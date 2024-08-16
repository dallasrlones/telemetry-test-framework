package logger

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"

    // "telemetry-test-framework/pkg/sqs"
)

// LogTelemetry logs the given telemetry data to a JSON file with a UUID and timestamp in the filename
func LogTelemetry(activity string, uuid string, startTime string, data interface{}) error {
    logDir := fmt.Sprintf("./telemetry-logs/%s", startTime)
    err := os.MkdirAll(logDir, os.ModePerm)
    if err != nil {
        return fmt.Errorf("failed to create log directory: %v", err)
    }

    fileName := fmt.Sprintf("%s_%s_%s.json", time.Now().UTC().Format("20060102150405"), uuid, activity)
    filePath := filepath.Join(logDir, fileName)

    file, err := os.Create(filePath)
    if err != nil {
        return fmt.Errorf("failed to create log file: %v", err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    err = encoder.Encode(data)
    if err != nil {
        return fmt.Errorf("failed to write log entry: %v", err)
    }

    // Send the telemetry data to SQS
    // sqsService := sqs.NewSQSService() // Get the SQS service instance
    // err = sqsService.SendMessage(data)
    // if err != nil {
    //     return fmt.Errorf("failed to send telemetry data to SQS: %v", err)
    // }

    return nil
}

// LogMessage logs a formatted message using log.Println and writes the same message to a file with the UUID as the filename
func LogMessage(uuid string, startTime string, format string, args ...interface{}) error {
    logDir := fmt.Sprintf("./telemetry-logs/%s", startTime)
    err := os.MkdirAll(logDir, os.ModePerm)
    if err != nil {
        return fmt.Errorf("failed to create log directory: %v", err)
    }

    fileName := fmt.Sprintf("%s_%s.log", startTime, uuid)
    filePath := filepath.Join(logDir, fileName)

    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to open log file: %v", err)
    }
    defer file.Close()

    // Format the message using fmt.Sprintf
    message := fmt.Sprintf(format, args...)

    // Log to the standard logger
    log.Println(message)

    // Also write the formatted log message to the file
    _, err = file.WriteString(fmt.Sprintf("%s: %s\n", time.Now().Format(time.RFC3339), message))
    if err != nil {
        return fmt.Errorf("failed to write log message to file: %v", err)
    }

    return nil
}
