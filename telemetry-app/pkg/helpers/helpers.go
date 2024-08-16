package helpers

import (
    "os"
    "os/signal"
    "strings"
    "syscall"
	"telemetry-test-framework/pkg/logger"
)

// Splits a comma-separated string of arguments into a slice of strings
func SplitArgs(argStr string) []string {
    return strings.Split(argStr, ",")
}

// Sets up graceful shutdown to save state before exiting
func SetupGracefulShutdown() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    go func() {
        <-c
        // Save the state and perform cleanup if necessary
        os.Exit(0)
    }()
}

func CheckAndSetEnvArgs(sessionUUID, startTime string, envName string) []string {
	execArgs := os.Getenv(envName)
    var argList []string
    if execArgs != "" {
        argList = SplitArgs(execArgs)
        logger.LogMessage(sessionUUID, startTime, "Executable arguments: %v\n", argList)
    }
	return argList
}

func CheckAndSetEnv(sessionUUID, startTime, envName string) string {
    envVal := os.Getenv(envName)
    if envVal == "" {
        logger.LogMessage(sessionUUID, startTime, "%s environment variable not set.", envName)
    } else {
        logger.LogMessage(sessionUUID, startTime, "%s set to %s\n", envName, envVal)
    }
    return envVal
}