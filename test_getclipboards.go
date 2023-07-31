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
			log.Printf("Received: %s", in)
		}
	}()

	getClipboardsResponse, err := client.GetClipboards(context.Background(), &ClipboardService.GetClipboardsRequest{Ids: []string{"1", "2"}})
	if err != nil {
		log.Fatalf("Error getting clipboards: %v", err)
	}
	for _, item := range getClipboardsResponse.Clipboards {
		log.Printf("Got clipboard: ID=%s, Content=%s", item.Id, item.Content)
	}
}
