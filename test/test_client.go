package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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
			log.Printf("Received: %s", in)
		}
	}()

	randomId, _ := uuid.NewUUID()

	// Call CreateClipboards every 5 seconds.
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		time.Sleep(1 * time.Second)
		createClipboardsResponse, err := client.CreateClipboards(context.Background(), &ClipboardService.CreateClipboardsRequest{Values: []string{"Hello", "World"}, DeviceId: randomId.String()})
		if err != nil {
			log.Fatalf("Error creating clipboards: %v", err)
		}
		log.Printf("Created clipboards with IDs: %v", createClipboardsResponse.Ids)
		time.Sleep(1 * time.Second)
		getClipboardsResponse, err := client.GetClipboards(context.Background(), &ClipboardService.GetClipboardsRequest{Ids: createClipboardsResponse.Ids})
		if err != nil {
			log.Fatalf("Error getting clipboards: %v", err)
		}
		for _, item := range getClipboardsResponse.Clipboards {
			log.Printf("Got clipboard: ID=%s, Content=%s, DeviceID=%s", item.Id, item.Content, item.DeviceId)
		}
		time.Sleep(1 * time.Second)
		mask := make([]*fieldmaskpb.FieldMask, 1)
		mask[0] = &fieldmaskpb.FieldMask{Paths: []string{"content"}}
		maskClipboardsResponse, err := client.GetClipboards(context.Background(), &ClipboardService.GetClipboardsRequest{Ids: []string{"1", "2"}, Mask: mask})
		if err != nil {
			log.Fatalf("Error getting clipboards: %v", err)
		}
		for _, item := range maskClipboardsResponse.Clipboards {
			log.Printf("Got clipboard: ID=%s, Content=%s", item.Id, item.Content)
		}
		time.Sleep(1 * time.Second)
		mask = make([]*fieldmaskpb.FieldMask, 1)
		mask[0] = &fieldmaskpb.FieldMask{Paths: []string{"device_id"}}
		maskClipboardsResponse, err = client.GetClipboards(context.Background(), &ClipboardService.GetClipboardsRequest{Ids: []string{"1", "2"}, Mask: mask})
		if err != nil {
			log.Fatalf("Error getting clipboards: %v", err)
		}
		for _, item := range maskClipboardsResponse.Clipboards {
			log.Printf("Got clipboard: ID=%s, deviceID=%s", item.Id, item.DeviceId)
		}
		time.Sleep(1 * time.Second)
		deleteClipboardsResponse, err := client.DeleteClipboards(context.Background(), &ClipboardService.DeleteClipboardsRequest{Ids: createClipboardsResponse.Ids, DeviceId: randomId.String()})
		if err != nil {
			log.Fatalf("Error deleting clipboards: %v", err)
		}
		if deleteClipboardsResponse.Success {
			log.Printf("Deleted clipboards: %v", createClipboardsResponse.Ids)
		} else {
			log.Printf("Failed to delete clipboards: %v", createClipboardsResponse.Ids)
		}
		time.Sleep(1 * time.Second)
		updateResponse, err := client.Update(context.Background(), &ClipboardService.UpdateRequest{Id: createClipboardsResponse.Ids[0], NewContent: "Hello World", DeviceId: randomId.String()})
		if err != nil {
			log.Fatalf("Error updating clipboard: %v", err)
		}
		if updateResponse.Success {
			log.Printf("Updated clipboard: %v", createClipboardsResponse.Ids[0])
		} else {
			log.Printf("Failed to update clipboard: %v", createClipboardsResponse.Ids[0])
		}
	}
}
