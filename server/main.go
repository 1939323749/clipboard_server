package main

import (
	"context"
	"encoding/json"
	"fmt"
	ClipboardService "github.com/1939323749/clipboard_server/proto"
	"github.com/1939323749/clipboard_server/server/message"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"strings"
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
	DeviceID  string    `bson:"deviceID"`
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

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	}

	s := grpc.NewServer(opts...)
	ClipboardService.RegisterClipboardServiceServer(s, &server{db: mongoClient.Database("strings")})

	reflection.Register(s)

	log.Printf("serving on %v", lis.Addr())

	go func() {
		RESTfulServer()
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("unary interceptor: %v", info.FullMethod)
	switch info.FullMethod {
	case "/proto.ClipboardService/CreateClipboards":
		return handler(ctx, req)
	}
	return handler(ctx, req)
}

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("stream interceptor: %v", info.FullMethod)
	return handler(srv, ss)
}

func (s *server) CreateClipboards(ctx context.Context, in *ClipboardService.CreateClipboardsRequest) (*ClipboardService.CreateClipboardsResponse, error) {
	operation := message.NewSendAble(in.Values[0])
	if _, ok := operation.(*message.Command); ok {
		return s.commandHandler(ctx, in)
	}
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

func (s *server) commandHandler(ctx context.Context, in *ClipboardService.CreateClipboardsRequest) (*ClipboardService.CreateClipboardsResponse, error) {
	command := in.Values[0][1:strings.IndexAny(in.Values[0], " ")]
	switch command {
	case "test":
		return &ClipboardService.CreateClipboardsResponse{Ids: []string{"test"}}, nil
	case "hello":
		return &ClipboardService.CreateClipboardsResponse{Ids: []string{"hello" + in.Values[0][strings.IndexAny(in.Values[0], " ")+1:]}}, nil
	case "ai":
		return s.aiHandler(ctx, in)
	default:
		return &ClipboardService.CreateClipboardsResponse{Ids: []string{"sorry, I don't know this command"}}, nil
	}
}

type Response struct {
	Reply string `json:"reply"`
	CID   string `json:"cid"`
}

// aiHandler is a handler for ai command
func (s *server) aiHandler(ctx context.Context, in *ClipboardService.CreateClipboardsRequest) (*ClipboardService.CreateClipboardsResponse, error) {
	decodedValue := url.QueryEscape(in.Values[0][strings.IndexAny(in.Values[0], " ")+1:])
	// replace this with your own ai service, it should return a json string like this:
	// {"reply": "hello", "cid": "test"}
	// I am using aircode for testing
	u := os.Getenv("CHAT_AI_URL") + decodedValue

	resp, err := http.Get(u)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}
	id := uuid.New().String()
	_, err = s.db.Collection(collection).InsertOne(ctx, &StringItem{ID: id, Value: in.Values[0], CreatedAt: time.Now(), DeviceID: in.DeviceId})
	if err != nil {
		return nil, err
	}
	for _, subscriber := range s.subscribers {
		if err := subscriber.Send(&ClipboardService.ClipboardMessage{Items: []*ClipboardService.ClipboardItem{{Id: id, Content: in.Values[0], DeviceId: in.DeviceId}}, Operation: "create"}); err != nil {
			log.Printf("Failed to send string to subscriber: %v", err)
		}
	}
	_, err = s.db.Collection(collection).InsertOne(ctx, &StringItem{ID: response.CID, Value: response.Reply, CreatedAt: time.Now(), DeviceID: "ai"})
	if err != nil {
		return nil, err
	}
	for _, subscriber := range s.subscribers {
		if err := subscriber.Send(&ClipboardService.ClipboardMessage{Items: []*ClipboardService.ClipboardItem{{Id: response.CID, Content: response.Reply, DeviceId: "ai"}}, Operation: "create"}); err != nil {
			log.Printf("Failed to send string to subscriber: %v", err)
		}
	}
	return &ClipboardService.CreateClipboardsResponse{Ids: []string{response.CID}, Msg: []string{response.Reply}}, nil
}

func (s *server) GetClipboards(ctx context.Context, in *ClipboardService.GetClipboardsRequest) (*ClipboardService.GetClipboardsResponse, error) {
	clipboardCollection := s.db.Collection(collection)
	var clipboardItems []*ClipboardService.ClipboardItem

	opts := options.Find().SetSort(bson.D{{"createdAt", -1}})
	cursor, err := clipboardCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Printf("Error closing cursor: %v", err)
		}
	}(cursor, ctx)

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

func (s *server) SubscribeClipboard(_ *ClipboardService.SubscribeClipboardRequest, stream ClipboardService.ClipboardService_SubscribeClipboardServer) error {
	s.subscribers = append(s.subscribers, stream)
	<-stream.Context().Done()
	for i, subscriber := range s.subscribers {
		if subscriber == stream {
			copy(s.subscribers[i:], s.subscribers[i+1:])
			s.subscribers = s.subscribers[:len(s.subscribers)-1]
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
			if err := subscriber.Send(&ClipboardService.ClipboardMessage{Items: []*ClipboardService.ClipboardItem{{Id: id, DeviceId: in.DeviceId}}, Operation: "delete"}); err != nil {
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
			log.Printf("error receiving from stream: %v", err)
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

func (s *server) CheckConnectivity(stream ClipboardService.ClipboardService_CheckConnectivityServer) error {
	for {
		_, err := stream.Recv()
		if err != nil {
			log.Printf("error receiving from stream: %v", err)
			return nil
		}
		err = stream.Send(&ClipboardService.Alive{})
		if err != nil {
			return nil
		}
	}
}

func (s *server) Update(ctx context.Context, in *ClipboardService.UpdateRequest) (*ClipboardService.UpdateResponse, error) {
	stringCollection := s.db.Collection(collection)
	log.Printf("Updating %v", in.Id)
	log.Printf("Updating %v", in.NewContent)
	_, err := stringCollection.UpdateOne(ctx, bson.M{"_id": in.Id}, bson.M{"$set": bson.M{"Value": in.NewContent}})
	if err != nil {
		return &ClipboardService.UpdateResponse{Success: false}, err
	}
	for _, subscriber := range s.subscribers {
		if err := subscriber.Send(&ClipboardService.ClipboardMessage{Items: []*ClipboardService.ClipboardItem{{Id: in.Id, Content: in.NewContent, DeviceId: in.DeviceId}}, Operation: "update"}); err != nil {
			log.Printf("Failed to send string to subscriber: %v", err)
		}
	}
	return &ClipboardService.UpdateResponse{Success: true}, nil
}

func RESTfulServer() {
	const token = "clipboard"
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	var client = ClipboardService.NewClipboardServiceClient(conn)
	http.HandleFunc("/api/create", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			if request.Header.Get("Content-Type") != "application/json" {
				http.Error(writer, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
				return
			}
			if request.Header.Get("Token") != token {
				http.Error(writer, "Invalid token", http.StatusUnauthorized)
				return
			}
			var clipboardItem ClipboardService.ClipboardItem
			err := json.NewDecoder(request.Body).Decode(&clipboardItem)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
			_, err = client.CreateClipboards(context.Background(), &ClipboardService.CreateClipboardsRequest{Values: []string{clipboardItem.Content}, DeviceId: clipboardItem.DeviceId})
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
			writer.WriteHeader(http.StatusCreated)
		} else {
			http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
	log.Printf("Starting server on port 50052")
	log.Fatal(http.ListenAndServe(":50052", nil))
}
