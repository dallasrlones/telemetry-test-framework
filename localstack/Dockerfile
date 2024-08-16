FROM localstack/localstack:latest

COPY create-queue.sh /usr/local/bin/create-queue.sh
RUN chmod +x /usr/local/bin/create-queue.sh

# Run LocalStack and the script
CMD ["sh", "-c", "localstack start -d && /usr/local/bin/create-queue.sh && localstack wait-for-services sqs"]
