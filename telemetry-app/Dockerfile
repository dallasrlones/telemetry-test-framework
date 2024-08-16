FROM golang:1.21.13-alpine

# Install git
RUN apk add --no-cache git

# Set Go cache directory
ENV GOCACHE=/go/cache

# Create the cache directory with appropriate permissions
RUN mkdir -p /go/cache && chmod -R 777 /go/cache

# Set working directory
WORKDIR /go/src/telemetry-test-framework

# Copy source code
COPY . .

RUN go mod tidy

# Copy the entrypoint script into the container
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Set entrypoint to the script
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
