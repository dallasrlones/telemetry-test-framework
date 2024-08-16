package sqs

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "sync"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSService struct {
    client     *sqs.Client
    queueURL   string
    sync.Mutex // Ensures that only one goroutine can send a message at a time
}

// Global instance of the SQSService
var instance *SQSService
var once sync.Once

// NewSQSService initializes the SQS service with a singleton pattern
func NewSQSService() *SQSService {
    once.Do(func() {
        // Load AWS configuration
        cfg, err := config.LoadDefaultConfig(context.TODO(),
            config.WithRegion(os.Getenv("AWS_REGION")),
            config.WithEndpointResolver(aws.EndpointResolverFunc(
                func(service, region string) (aws.Endpoint, error) {
                    if service == sqs.ServiceID {
                        endpoint := os.Getenv("SQS_ENDPOINT")
                        log.Printf("Resolved SQS endpoint: %s", endpoint)
                        return aws.Endpoint{
                            URL: endpoint,
                        }, nil
                    }
                    return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested for service %s", service)
                })),
        )
        if err != nil {
            log.Fatalf("unable to load SDK config, %v", err)
        }

        // Get the queue URL from environment variables
        queueURL := os.Getenv("SQS_QUEUE_URL")
        if queueURL == "" {
            log.Fatalf("SQS_QUEUE_URL environment variable is not set")
        }

        log.Printf("Using SQS Queue URL: %s", queueURL)

        // Initialize the SQSService instance
        instance = &SQSService{
            client:   sqs.NewFromConfig(cfg),
            queueURL: queueURL,
        }
    })

    return instance
}

// SendMessage sends a message to the SQS queue
func (s *SQSService) SendMessage(message interface{}) error {
    s.Lock()
    defer s.Unlock()

    // Marshal the message into JSON
    messageBody, err := json.Marshal(message)
    if err != nil {
        return fmt.Errorf("failed to marshal message: %w", err)
    }

    // Send the message to the SQS queue
    _, err = s.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
        QueueUrl:    &s.queueURL,
        MessageBody: aws.String(string(messageBody)),
    })
    if err != nil {
        return fmt.Errorf("failed to send message to SQS: %w", err)
    }

    return nil
}
