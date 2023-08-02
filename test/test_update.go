package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/proto"
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
		updateResponse, err := client.Update(context.Background(), &ClipboardService.UpdateRequest{Id: createClipboardsResponse.Ids[index], NewContent: "World"})

		if err != nil {
			log.Fatalf("Error updating clipboard: %v", err)
		}
		if updateResponse.Success {
			log.Printf("Updated clipboard with ID: %s", createClipboardsResponse.Ids[index])
		}
	}
}
