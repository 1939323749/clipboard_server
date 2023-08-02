package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/proto"
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

	stream, err := client.StreamMessage(context.Background())
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
	ticker := time.NewTicker(5 * time.Second)
	println("Sending")
	for range ticker.C {
		err := stream.Send(&ClipboardService.StreamMsg{Msg: "Hello"})
		if err != nil {
			log.Fatalf("Error sending: %v", err)
		}
	}
	select {}
}
