#!/bin/bash

echo "LocalStack is up, creating SQS queue..."

# Create the SQS queue
awslocal sqs create-queue --queue-name test-queue

if [ $? -eq 0 ]; then
  echo "SQS queue created successfully."
else
  echo "Failed to create SQS queue."
fi

# Keep the container running
exec "$@"
