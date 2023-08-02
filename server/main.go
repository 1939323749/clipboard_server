package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/proto"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	DeviceID  string    `bson:"deviceID,omitempty"`
}

var collection = "clipboards"

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

	log.Printf("serving on %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CreateClipboards(ctx context.Context, in *ClipboardService.CreateClipboardsRequest) (*ClipboardService.CreateClipboardsResponse, error) {
	stringCollection := s.db.Collection(collection)
	var ids []string
	for _, value := range in.Values {
		id := uuid.New().String()
		ids = append(ids, id)
		_, err := stringCollection.InsertOne(ctx, &StringItem{ID: id, Value: value, CreatedAt: time.Now(), DeviceID: in.DeviceId})
		if err != nil {
			return nil, err
		}
		for _, subscriber := range s.subscribers {
			if err := subscriber.Send(&ClipboardService.ClipboardMessage{Items: []*ClipboardService.ClipboardItem{{Id: id, Content: value, DeviceId: in.DeviceId}}, Operation: "create"}); err != nil {
				log.Printf("Failed to send string to subscriber: %v", err)
			}
		}
	}
	return &ClipboardService.CreateClipboardsResponse{Ids: ids}, nil
}

func (s *server) GetClipboards(ctx context.Context, in *ClipboardService.GetClipboardsRequest) (*ClipboardService.GetClipboardsResponse, error) {
	clipboardCollection := s.db.Collection(collection)
	var clipboardItems []*ClipboardService.ClipboardItem

	opts := options.Find().SetSort(bson.D{{"createdAt", -1}})
	cursor, err := clipboardCollection.Find(ctx, bson.M{}, opts)
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
		if in.GetMask() != nil {
			clipboardItem := &ClipboardService.ClipboardItem{}
			fm := in.GetMask()[0]
			if fm.GetPaths()[0] == "content" {
				clipboardItem.Content = item.Value
			}
			if fm.GetPaths()[0] == "device_id" {
				clipboardItem.DeviceId = item.DeviceID
			}
			clipboardItem.Id = item.ID
			clipboardItems = append(clipboardItems, clipboardItem)
		} else {
			clipboardItems = append(clipboardItems, &ClipboardService.ClipboardItem{Id: item.ID, Content: item.Value, DeviceId: item.DeviceID})
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &ClipboardService.GetClipboardsResponse{Clipboards: clipboardItems}, nil
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

func (s *server) DeleteClipboards(ctx context.Context, in *ClipboardService.DeleteClipboardsRequest) (*ClipboardService.DeleteClipboardsResponse, error) {
	stringCollection := s.db.Collection(collection)
	for _, id := range in.Ids {
		_, err := stringCollection.DeleteOne(ctx, bson.M{"_id": id})
		if err != nil {
			return &ClipboardService.DeleteClipboardsResponse{Success: false}, err
		}
		for _, subscriber := range s.subscribers {
			if err := subscriber.Send(&ClipboardService.ClipboardMessage{Items: []*ClipboardService.ClipboardItem{{Id: id}}, Operation: "delete"}); err != nil {
				log.Printf("Failed to send string to subscriber: %v", err)
			}
		}
	}
	return &ClipboardService.DeleteClipboardsResponse{Success: true}, nil
}

func (s *server) StreamMessage(stream ClipboardService.ClipboardService_StreamMessageServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		err = stream.Send(&ClipboardService.StreamMsg{Msg: "Stream beginning!", Timestamp: timestamppb.New(time.Now())})
		for i := 0; i < 10; i++ {
			err = stream.Send(&ClipboardService.StreamMsg{Msg: "hello " + req.Msg, Timestamp: timestamppb.New(time.Now())})
			if err != nil {
				return err
			}
		}
	}
}

func (s *server) Update(ctx context.Context, in *ClipboardService.UpdateRequest) (*ClipboardService.UpdateResponse, error) {
	stringCollection := s.db.Collection(collection)
	_, err := stringCollection.UpdateOne(ctx, bson.M{"_id": in.Id}, bson.M{"$set": bson.M{"value": in.NewContent}})
	if err != nil {
		return &ClipboardService.UpdateResponse{Success: false}, err
	}
	for _, subscriber := range s.subscribers {
		if err := subscriber.Send(&ClipboardService.ClipboardMessage{Items: []*ClipboardService.ClipboardItem{{Id: in.Id, Content: in.NewContent}}, Operation: "update"}); err != nil {
			log.Printf("Failed to send string to subscriber: %v", err)
		}
	}
	return &ClipboardService.UpdateResponse{Success: true}, nil
}
