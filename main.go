package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/clipboard_service"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "net/http/pprof"
)

type server struct {
	ClipboardService.UnimplementedClipboardServiceServer
	db          *mongo.Database
	subscribers []ClipboardService.ClipboardService_SubscribeClipboardServer
}

type StringItem struct {
	ID    string `bson:"_id,omitempty"`
	Value string `bson:"value,omitempty"`
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	ClipboardService.RegisterClipboardServiceServer(s, &server{db: mongoClient.Database("strings")})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CreateClipboards(ctx context.Context, in *ClipboardService.CreateClipboardsRequest) (*ClipboardService.CreateClipboardsResponse, error) {
	stringCollection := s.db.Collection("stringCollection")
	var ids []string
	for _, value := range in.Values {
		id := uuid.New().String()
		ids = append(ids, id)
		_, err := stringCollection.InsertOne(ctx, &StringItem{ID: id, Value: value})
		if err != nil {
			return nil, err
		}
		for _, subscriber := range s.subscribers {
			if err := subscriber.Send(&ClipboardService.ClipboardMessage{Value: value}); err != nil {
				log.Printf("Failed to send string to subscriber: %v", err)
			}
		}
	}
	return &ClipboardService.CreateClipboardsResponse{Ids: ids}, nil
}

func (s *server) GetClipboards(ctx context.Context, in *ClipboardService.GetClipboardsRequest) (*ClipboardService.GetClipboardsResponse, error) {
	stringCollection := s.db.Collection("stringCollection")
	var values []string
	for _, id := range in.Ids {
		result := stringCollection.FindOne(ctx, &StringItem{ID: id})
		if result.Err() != nil {
			return nil, result.Err()
		}
		var item StringItem
		err := result.Decode(&item)
		if err != nil {
			return nil, err
		}
		values = append(values, item.Value)
	}
	return &ClipboardService.GetClipboardsResponse{Values: values}, nil
}

func (s *server) SubscribeClipboard(in *ClipboardService.SubscribeClipboardRequest, stream ClipboardService.ClipboardService_SubscribeClipboardServer) error {
	s.subscribers = append(s.subscribers, stream)
	<-stream.Context().Done()

	for i, subscriber := range s.subscribers {
		if subscriber == stream {
			s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
			break
		}
	}
	return nil
}
