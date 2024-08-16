#!/bin/sh

# Set environment variables with hardcoded paths
export GO_ENV=development
export EXEC_PATH=/go/src/telemetry-test-framework/scripts/echo_script.sh
export EXEC_ARGS=${EXEC_ARGS:---default-flag=value}

# Hardcode the file paths similar to EXEC_PATH
export FILE_CREATE_PATH=/go/src/telemetry-test-framework/test_file.txt
export FILE_CONTENT=${FILE_CONTENT:-"Test content"}
export FILE_UPDATE_PATH=/go/src/telemetry-test-framework/test_file.txt
export FILE_MODIFY_CONTENT=${FILE_MODIFY_CONTENT:-/go/src/telemetry-test-framework/test_file.txt}
export FILE_DELETE_PATH=${FILE_DELETE_PATH:-/go/src/telemetry-test-framework/test_file.txt}
export HTTP_ENDPOINT=${HTTP_ENDPOINT:-"node-server"}
export HTTP_PORT=3000
export AWS_REGION=us-east-1
export AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID:-test}
export AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY:-test}
export SQS_QUEUE_URL=${SQS_QUEUE_URL:-http://localstack:4566/queue/test-queue}
export SQS_ENDPOINT=${SQS_ENDPOINT:-http://localstack:4566}
export USER_ID=$(id -u)
export GROUP_ID=$(id -g)
export USER=$(whoami)

# Run Docker Compose without rebuilding
docker compose up
