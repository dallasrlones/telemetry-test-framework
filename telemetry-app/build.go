package main

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "log"
)

func main() {
    fmt.Println("Starting build process...")

    outputDir := os.Getenv("BUILD_LOC")
    if outputDir == "" {
        log.Fatalf("BUILD_LOC environment variable not set.")
    }
    outputName := "telemetry-test-framework"

    if runtime.GOOS == "windows" {
        outputName += ".exe"
    }

    // Run the go build command
    cmd := exec.Command("go", "build", "-o", outputDir+outputName, "./cmd/telemetry")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        fmt.Printf("Build failed: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Build completed successfully.")
}
