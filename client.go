package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/clipboard_service"
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

	stream, err := client.SubscribeClipboard(context.Background(), &ClipboardService.SubscribeClipboardRequest{})
	if err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}

	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving: %v", err)
			}
			log.Printf("Received: %s", in.Value)
		}
	}()

	// Call CreateClipboards every 5 seconds.
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		createClipboardsResponse, err := client.CreateClipboards(context.Background(), &ClipboardService.CreateClipboardsRequest{Values: []string{"Hello", "World"}})
		if err != nil {
			log.Fatalf("Error creating clipboards: %v", err)
		}
		log.Printf("Created clipboards with IDs: %v", createClipboardsResponse.Ids)

		getClipboardsResponse, err := client.GetClipboards(context.Background(), &ClipboardService.GetClipboardsRequest{Ids: createClipboardsResponse.Ids})
		if err != nil {
			log.Fatalf("Error getting clipboards: %v", err)
		}
		for _, item := range getClipboardsResponse.Clipboards {
			log.Printf("Got clipboard: ID=%s, Content=%s", item.Id, item.Content)
		}

		deleteClipboardsResponse, err := client.DeleteClipboards(context.Background(), &ClipboardService.DeleteClipboardsRequest{Ids: createClipboardsResponse.Ids})
		if err != nil {
			log.Fatalf("Error deleting clipboards: %v", err)
		}
		if deleteClipboardsResponse.Success {
			log.Printf("Deleted clipboards: %v", createClipboardsResponse.Ids)
		} else {
			log.Printf("Failed to delete clipboards: %v", createClipboardsResponse.Ids)
		}
	}
}
