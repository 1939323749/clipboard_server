package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/clipboard_service"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	_ "net/http/pprof"
	"time"
)

type server struct {
	ClipboardService.UnimplementedClipboardServiceServer
	db          *mongo.Database
	subscribers []ClipboardService.ClipboardService_SubscribeClipboardServer
}

type StringItem struct {
	ID        string    `bson:"_id,omitempty"`
	Value     string    `bson:"value,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty"`
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

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CreateClipboards(ctx context.Context, in *ClipboardService.CreateClipboardsRequest) (*ClipboardService.CreateClipboardsResponse, error) {
	stringCollection := s.db.Collection("strings")
	var ids []string
	for _, value := range in.Values {
		id := uuid.New().String()
		ids = append(ids, id)
		_, err := stringCollection.InsertOne(ctx, &StringItem{ID: id, Value: value, CreatedAt: time.Now()})
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
	stringCollection := s.db.Collection("strings")
	var values []string

	opts := options.Find().SetSort(bson.D{{"createdAt", -1}})
	cursor, err := stringCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item StringItem
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		values = append(values, item.Value)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
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
