package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := ClipboardService.NewClipboardServiceClient(conn)

	randomId, _ := uuid.NewUUID()

	// Call CreateClipboards every 5 seconds.
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		createClipboardsResponse, err := client.CreateClipboards(context.Background(), &ClipboardService.CreateClipboardsRequest{Values: []string{"Hello"}, DeviceId: randomId.String()})
		if err != nil {
			log.Fatalf("Error creating clipboards: %v", err)
		}
		log.Printf("Created clipboards with IDs: %v", createClipboardsResponse.Ids)
	}
}
