package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error closing connection: %v", err)
		}
	}(conn)

	client := ClipboardService.NewClipboardServiceClient(conn)

	stream, err := client.CheckConnectivity(context.Background())
	if err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}

	go func() {
		for {
			_, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving: %v", err)
			}
			log.Printf("Alive")
		}
	}()
	ticker := time.NewTicker(1 * time.Second)
	println("Sending")
	for range ticker.C {
		err := stream.Send(&ClipboardService.Alive{})
		if err != nil {
			log.Fatalf("Error sending: %v", err)
		}
	}
	select {}
}
