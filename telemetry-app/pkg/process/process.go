package process

import (
    "fmt"
    "os"
    "os/exec"
    "runtime" // Import the runtime package to get OS details
    "telemetry-test-framework/pkg/logger"
    "telemetry-test-framework/pkg/helpers"
    "time"
)

// ProcessInfo holds details about the process that was created
type ProcessInfo struct {
    Timestamp   time.Time `json:"timestamp"`
    Username    string    `json:"username"`
    ProcessName string    `json:"process_name"`
    CommandLine string    `json:"command_line"`
    ProcessID   int       `json:"process_id"`
    OS          string    `json:"os"` // Add OS field to capture the operating system
}

// getUsername retrieves the current user's username in a cross-platform way.
func getUsername() (string, error) {
    switch runtime.GOOS {
    case "windows":
        return os.Getenv("USERNAME"), nil
    case "darwin", "linux":
        return os.Getenv("USER"), nil
    default:
        return "", fmt.Errorf("unsupported platform: %s", runtime.GOOS)
    }
}

// RunProcessCreation starts a process and logs the telemetry data.
func runProcessCreation(uuid, startTime, execPath string, args []string) (ProcessInfo, error) {
    cmd := exec.Command(execPath, args...)
    err := cmd.Start()
    if err != nil {
        return ProcessInfo{}, fmt.Errorf("failed to start process: %v", err)
    }

    // Capture process-related information
    processName := execPath
    processID := cmd.Process.Pid
    commandLine := fmt.Sprintf("%s %v", execPath, args)

    username, err := getUsername()
    if err != nil {
        return ProcessInfo{}, fmt.Errorf("failed to retrieve username: %v", err)
    }

    processInfo := ProcessInfo{
        Timestamp:   time.Now(),
        Username:    username,
        ProcessName: processName,
        CommandLine: commandLine,
        ProcessID:   processID,
        OS:          runtime.GOOS, // Capture the current operating system
    }

    // Log the process creation telemetry with the UUID
    err = logger.LogTelemetry("process_creation", uuid, startTime, processInfo)
    if err != nil {
        return ProcessInfo{}, fmt.Errorf("failed to log process creation telemetry: %v", err)
    }

    return processInfo, nil
}

func RunProcessOperation(sessionUUID string, startTime string, errors *[]error) {
    execPath := helpers.CheckAndSetEnv(sessionUUID, startTime, "EXEC_PATH")
    argList := helpers.CheckAndSetEnvArgs(sessionUUID, startTime, "EXEC_ARGS")

    logger.LogMessage(sessionUUID, startTime, "Running process creation operation...")
    _, err := runProcessCreation(sessionUUID, startTime, execPath, argList)
    if err != nil {
        logger.LogMessage(sessionUUID, startTime, "Process creation failed: %v\n", err)
        *errors = append(*errors, err)
    } else {
        logger.LogMessage(sessionUUID, startTime, "Process creation succeeded.")
    }
}