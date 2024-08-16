#!/bin/sh

# Define output directory based on BUILD_LOC or use default
outputDir=${BUILD_LOC:-/go/src/telemetry-test-framework}
outputName="telemetry-test-framework"

if [ -n "$BUILD_LOC" ]; then
    # If BUILD_LOC is set, build the binary
    echo "Building binary for OS=${GOOS} ARCH=${GOARCH}..."

    go build -o "$outputDir/$outputName" ./cmd/telemetry

    if [ $? -eq 0 ]; then
        echo "Build complete. Binary created at $outputDir/$outputName."
    else
        echo "Build failed."
        exit 1
    fi
else
    # If BUILD_LOC is not set, run the main.go file directly
    echo "Running main.go..."

    go run ./cmd/telemetry/main.go
fi
