package main

import (
    "os"
    "time"
    "github.com/google/uuid"
    "telemetry-test-framework/pkg/logger"
    "telemetry-test-framework/pkg/helpers"
    "telemetry-test-framework/pkg/process"
    "telemetry-test-framework/pkg/fileops"
    "telemetry-test-framework/pkg/network"
)

func main() {
    helpers.SetupGracefulShutdown()
    var errors []error
    sessionUUID := uuid.New().String()
    startTime := time.Now().UTC().Format("20060102150405")
    logger.LogMessage(sessionUUID, startTime, "Starting telemetry test framework...Generated session UUID: %s at %s\n", sessionUUID, startTime)

    process.RunProcessOperation(sessionUUID, startTime, &errors)
    // do we want read too?
    fileops.RunFileCreateOperation(sessionUUID, startTime, &errors)
    fileops.RunFileModifyOperation(sessionUUID, startTime, &errors)
    fileops.RunFileDeleteOperation(sessionUUID, startTime, &errors)
    
    network.RunNetworkOperation(sessionUUID, startTime, &errors)

    checkForErrors(sessionUUID, startTime, errors)
    success(sessionUUID, startTime)
}

// checkForErrors checks the error slice and handles any errors found
func checkForErrors(sessionUUID, startTime string, errors []error) {
    if len(errors) > 0 {
        logger.LogMessage(sessionUUID, startTime, "\nEncountered the following errors:\n")
        for _, err := range errors {
            logger.LogMessage(sessionUUID, startTime, "error: %v\n", err)
        }
        network.ShutDownTestingServer()
        os.Exit(1)
    }
}

func success(sessionUUID, startTime string) {
	logger.LogMessage(sessionUUID, startTime, "All operations completed successfully.")
    network.ShutDownTestingServerSuccess()
    os.Exit(0)
}