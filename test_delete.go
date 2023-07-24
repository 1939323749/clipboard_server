package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/clipboard_service"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := ClipboardService.NewClipboardServiceClient(conn)

	createClipboardsResponse, err := client.CreateClipboards(context.Background(), &ClipboardService.CreateClipboardsRequest{Values: []string{"Hello"}})
	if err != nil {
		log.Fatalf("Error creating clipboards: %v", err)
	}
	log.Printf("Created clipboards with IDs: %v", createClipboardsResponse.Ids)
	if len(createClipboardsResponse.Ids) > 0 {
		index := 0
		deleteClipboardsResponse, err := client.DeleteClipboards(
			context.Background(), &ClipboardService.DeleteClipboardsRequest{Ids: []string{createClipboardsResponse.Ids[index]}})
		if err != nil {
			log.Fatalf("Error deleting clipboard: %v", err)
		}
		if deleteClipboardsResponse.Success {
			log.Printf("Deleted clipboard with ID: %s", createClipboardsResponse.Ids[index])
		}
	}
}
