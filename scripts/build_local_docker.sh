#!/bin/sh

# Detect the OS and set GOOS and GOARCH accordingly
case "$(uname -s)" in
    Linux*)
        export GOOS=linux
        export GOARCH=amd64
        ;;
    Darwin*)
        export GOOS=darwin
        export GOARCH=amd64
        ;;
    *)
        echo "Unsupported OS detected."
        exit 1
        ;;
esac

# Print the detected OS and architecture
echo "Detected OS: $GOOS"
echo "Detected ARCH: $GOARCH"

# Run Docker Compose with the specified environment variables
# Pass all necessary environment variables to Docker Compose
GOOS="$GOOS" GOARCH="$GOARCH" BUILD_LOC="${BUILD_LOC:-./cmd/telemetry}" docker compose up --build
