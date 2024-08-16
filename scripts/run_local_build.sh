#!/bin/bash

# Get the directory of the current script
SCRIPT_DIR=$(dirname "$(realpath "$0")")

# Define the root path based on the location of the script
ROOT_DIR=$(realpath "$SCRIPT_DIR/..")

# Export environment variables
export GO_ENV=development
export EXEC_PATH="$ROOT_DIR/scripts/echo_script.sh"
# /go/src/telemetry-test-framework/scripts/echo_script.sh
export FILE_CREATE_PATH="$ROOT_DIR/test_file.txt"
export FILE_CONTENT=${FILE_CONTENT:-"Test content"}
export FILE_UPDATE_PATH="$ROOT_DIR/test_file.txt"
export FILE_MODIFY_CONTENT=${FILE_MODIFY_CONTENT:-"Modified content"}
export FILE_DELETE_PATH="$ROOT_DIR/test_file.txt"
export HTTP_ENDPOINT=${HTTP_ENDPOINT:-"google.com"}
export HTTP_PORT=80
export AWS_REGION=us-east-1
export AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID:-test}
export AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY:-test}
export SQS_QUEUE_URL=${SQS_QUEUE_URL:-http://localstack:4566/queue/test-queue}
export SQS_ENDPOINT=${SQS_ENDPOINT:-http://localstack:4566}
export USER_ID=$(id -u)
export GROUP_ID=$(id -g)
export USER=$(whoami)

# Path to the built binary relative to the root directory
BINARY_PATH="$ROOT_DIR/telemetry-app/cmd/telemetry/telemetry-test-framework"

# Verify if the binary exists and is executable
if [ -x "$BINARY_PATH" ]; then
  # Run the built binary
  $BINARY_PATH
else
  echo "Error: Binary file not found or not executable."
  exit 1
fi
