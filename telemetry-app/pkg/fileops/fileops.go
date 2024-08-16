package fileops

import (
    "os"
    "os/user"
    "log"
    "time"
    "fmt"
    "os/exec"
    "runtime"
    "telemetry-test-framework/pkg/logger"
    "telemetry-test-framework/pkg/helpers"
)

type FileActivity struct {
    Timestamp        time.Time `json:"timestamp"`
    FullPath         string    `json:"full_path"`
    Activity         string    `json:"activity"`
    Username         string    `json:"username"`
    ProcessName      string    `json:"process_name"`
    ProcessCmdLine   string    `json:"process_cmd_line"`
    ProcessID        int       `json:"process_id"`
    StartTime        string    `json:start_time`
}

// createFile creates a file with the specified content at the specified path
func createFile(uuid, startTime, filePath string, content string) (FileActivity, error) {
    log.Printf("Creating file at path: %s", filePath)
    err := os.WriteFile(filePath, []byte(content), 0644)
    if err != nil {
        log.Printf("Failed to create file at path: %s with error: %v", filePath, err)
        return FileActivity{}, err
    }

    activity, err := captureFileActivity(filePath, "create", startTime)
    if err != nil {
        log.Printf("Failed to capture file activity for create at path: %s with error: %v", filePath, err)
        return activity, err
    }

    // Log the file creation telemetry with the UUID
    err = logger.LogTelemetry("file_creation", uuid, startTime, activity)
    if err != nil {
        log.Printf("Failed to log telemetry for file creation at path: %s with error: %v", filePath, err)
        return activity, fmt.Errorf("failed to log file creation telemetry: %v", err)
    }

    log.Printf("Successfully created file and logged telemetry for path: %s", filePath)
    return activity, nil
}

// modifyFile modifies the content of an existing file at the specified path
func modifyFile(uuid, startTime, filePath string, newContent string) (FileActivity, error) {
    log.Printf("Modifying file at path: %s", filePath)
    err := os.WriteFile(filePath, []byte(newContent), 0644)
    if err != nil {
        log.Printf("Failed to modify file at path: %s with error: %v", filePath, err)
        return FileActivity{}, err
    }

    activity, err := captureFileActivity(filePath, "modify", startTime)
    if err != nil {
        log.Printf("Failed to capture file activity for modify at path: %s with error: %v", filePath, err)
        return activity, err
    }

    // Log the file modification telemetry with the UUID
    err = logger.LogTelemetry("file_modification", uuid, startTime, activity)
    if err != nil {
        log.Printf("Failed to log telemetry for file modification at path: %s with error: %v", filePath, err)
        return activity, fmt.Errorf("failed to log file modification telemetry: %v", err)
    }

    log.Printf("Successfully modified file and logged telemetry for path: %s", filePath)
    return activity, nil
}

// deleteFile deletes the file at the specified path
func deleteFile(uuid, startTime, filePath string) (FileActivity, error) {
    log.Printf("Deleting file at path: %s", filePath)
    err := os.Remove(filePath)
    if err != nil {
        log.Printf("Failed to delete file at path: %s with error: %v", filePath, err)
        return FileActivity{}, err
    }

    activity, err := captureFileActivity(filePath, "delete", startTime)
    if err != nil {
        log.Printf("Failed to capture file activity for delete at path: %s with error: %v", filePath, err)
        return activity, err
    }

    // Log the file deletion telemetry with the UUID
    err = logger.LogTelemetry("file_deletion", uuid, startTime, activity)
    if err != nil {
        log.Printf("Failed to log telemetry for file deletion at path: %s with error: %v", filePath, err)
        return activity, fmt.Errorf("failed to log file deletion telemetry: %v", err)
    }

    log.Printf("Successfully deleted file and logged telemetry for path: %s", filePath)
    return activity, nil
}

// captureFileActivity captures metadata about the file operation
func captureFileActivity(filePath, activity, startTime string) (FileActivity, error) {
    usr, err := user.Current()
    if err != nil {
        log.Printf("Failed to get current user for file operation at path: %s with error: %v", filePath, err)
        return FileActivity{}, err
    }

    pid := os.Getpid()
    cmdLine, err := getCommandLine(pid)
    if err != nil {
        log.Printf("Failed to get command line for file operation at path: %s with error: %v. Using fallback command line.", filePath, err)
        cmdLine = []byte(os.Args[0]) // Fallback to the executable name only if `ps` fails
    }

    procName := os.Args[0]

    return FileActivity{
        Timestamp:      time.Now(),
        FullPath:       filePath,
        Activity:       activity,
        Username:       usr.Username,
        ProcessName:    procName,
        ProcessCmdLine: string(cmdLine),
        ProcessID:      pid,
        StartTime:      startTime,
    }, nil
}

// getCommandLine returns the command line used to run the process based on the operating system
func getCommandLine(pid int) ([]byte, error) {
    switch runtime.GOOS {
    case "windows":
        return exec.Command("cmd", "/C", fmt.Sprintf("wmic process where ProcessId=%d get CommandLine", pid)).Output()
    case "darwin", "linux":
        return exec.Command("ps", "-p", fmt.Sprintf("%d", pid), "-o", "command=").Output()
    default:
        return nil, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
    }
}

func RunFileCreateOperation(sessionUUID, startTime string, errors *[]error) {
	fileCreatePath := helpers.CheckAndSetEnv(sessionUUID, startTime, "FILE_CREATE_PATH")
    fileContent := helpers.CheckAndSetEnv(sessionUUID, startTime, "FILE_CONTENT")

	logger.LogMessage(sessionUUID, startTime, "Running file creation operation...")
    if _, err := createFile(sessionUUID, startTime, fileCreatePath, fileContent); err != nil {
        logger.LogMessage(sessionUUID, startTime, "File creation failed: %v\n", err)
        *errors = append(*errors, err)
    } else {
        logger.LogMessage(sessionUUID, startTime, "File creation succeeded.")
    }
}

func RunFileModifyOperation(sessionUUID, startTime string, errors *[]error) {
	fileUpdatePath := helpers.CheckAndSetEnv(sessionUUID, startTime, "FILE_UPDATE_PATH")
    fileModifyContent := helpers.CheckAndSetEnv(sessionUUID, startTime, "FILE_MODIFY_CONTENT")

	logger.LogMessage(sessionUUID, startTime, "Running file modification operation...")
    if _, err := modifyFile(sessionUUID, startTime, fileUpdatePath, fileModifyContent); err != nil {
        logger.LogMessage(sessionUUID, startTime, "File modification failed: %v\n", err)
        *errors = append(*errors, err)
    } else {
        logger.LogMessage(sessionUUID, startTime, "File modification succeeded.")
    }
}

func RunFileDeleteOperation(sessionUUID, startTime string, errors *[]error) {
	fileDeletePath := helpers.CheckAndSetEnv(sessionUUID, startTime, "FILE_DELETE_PATH")

	logger.LogMessage(sessionUUID, startTime, "Running file deletion operation...")
    if _, err := deleteFile(sessionUUID, startTime, fileDeletePath); err != nil {
        logger.LogMessage(sessionUUID, startTime, "File deletion failed: %v\n", err)
        *errors = append(*errors, err)
    } else {
        logger.LogMessage(sessionUUID, startTime, "File deletion succeeded.")
    }
}