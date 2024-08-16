package network

import (
    "strconv"
    "fmt"
    "net"
    "net/http"
    "os/user"
    "syscall"
    "telemetry-test-framework/pkg/logger" // Import the logger package
    "telemetry-test-framework/pkg/helpers"
    "time"
)

// NetworkOperation holds details about the network activity that was performed
type NetworkOperation struct {
    Timestamp       time.Time `json:"timestamp"`
    Username        string    `json:"username"`
    DestinationAddr string    `json:"destination_address"`
    DestinationPort int       `json:"destination_port"`
    SourceAddr      string    `json:"source_address"`
    SourcePort      int       `json:"source_port"`
    DataSent        int64     `json:"data_sent"`
    Protocol        string    `json:"protocol"`
    ProcessName     string    `json:"process_name"`
    CommandLine     string    `json:"command_line"`
    ProcessID       int       `json:"process_id"`
    StartTime       string    `json:"start_time"`
}

// ShutDownTestingServer hits the node-server with a GET request to /quit
func ShutDownTestingServer() error {
    // Define the shutdown URL
    url := "http://node-server:3000/quit"

    // Perform the GET request
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("failed to send shutdown request: %v", err)
    }
    defer resp.Body.Close()

    return nil
}

// ShutdownTestingServerSuccess hits the node-server with a GET request to /success
func ShutDownTestingServerSuccess() error {
    // Define the shutdown URL
    url := "http://node-server:3000/success"

    // Perform the GET request
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("failed to send shutdown request: %v", err)
    }
    defer resp.Body.Close()

    return nil
}

// RunNetworkOperation establishes a network connection, sends data, and logs the telemetry data
func NetworkOperationCall(uuid, startTime, destinationAddr string, destinationPort int, data []byte) (NetworkOperation, error) {
    // Resolve the destination address
    destAddr := fmt.Sprintf("%s:%d", destinationAddr, destinationPort)
    conn, err := net.Dial("tcp", destAddr)
    if err != nil {
        return NetworkOperation{}, fmt.Errorf("failed to establish connection to %s: %v", destAddr, err)
    }
    defer conn.Close()

    // Capture source address and port
    localAddr := conn.LocalAddr().(*net.TCPAddr)

    // Send data
    bytesSent, err := conn.Write(data)
    if err != nil {
        return NetworkOperation{}, fmt.Errorf("failed to send data: %v", err)
    }

    // Capture process-related information
    processName := "network_operation" // Assuming this operation is running as a standalone process
    processID := syscall.Getpid()

    // Capture the username of the process starter
    user, err := user.Current()
    if err != nil {
        return NetworkOperation{}, fmt.Errorf("failed to get current user: %v", err)
    }

    // Build the network operation information
    networkOp := NetworkOperation{
        Timestamp:       time.Now(),
        StartTime:       startTime,
        Username:        user.Username,
        DestinationAddr: destinationAddr,
        DestinationPort: destinationPort,
        SourceAddr:      localAddr.IP.String(),
        SourcePort:      localAddr.Port,
        DataSent:        int64(bytesSent),
        Protocol:        "tcp",
        ProcessName:     processName,
        CommandLine:     fmt.Sprintf("connect to %s:%d and send data", destinationAddr, destinationPort),
        ProcessID:       processID,
    }

    // Log the network operation telemetry with the UUID
    err = logger.LogTelemetry("network_operation", uuid, startTime, networkOp)
    if err != nil {
        return NetworkOperation{}, fmt.Errorf("failed to log network operation telemetry: %v", err)
    }

    return networkOp, nil
}

func RunNetworkOperation(sessionUUID, startTime string, errors *[]error) {
    serverUrl := helpers.CheckAndSetEnv(sessionUUID, startTime, "HTTP_ENDPOINT")
    serverPortStr := helpers.CheckAndSetEnv(sessionUUID, startTime, "HTTP_PORT")
    serverPort, err := strconv.Atoi(serverPortStr)
    if err != nil {
        logger.LogMessage(sessionUUID, startTime, "Invalid port: %s, error: %v", serverPortStr, err)
        *errors = append(*errors, err)
        return
    }

    logger.LogMessage(sessionUUID, startTime, "Running network operation...")
    _, err = NetworkOperationCall(sessionUUID, startTime, serverUrl, serverPort, []byte("test data"))
    if err != nil {
        logger.LogMessage(sessionUUID, startTime, "Network operation failed: %v\n", err)
        *errors = append(*errors, err)
    } else {
        logger.LogMessage(sessionUUID, startTime, "Network operation succeeded.")
    }
}