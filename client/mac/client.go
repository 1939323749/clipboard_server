package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os/exec"
	"strings"
	"time"
)

var ignoreDeviceIdList = []string{"macOS_popclip"}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(retryInterceptor))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	client := ClipboardService.NewClipboardServiceClient(conn)

	stream, err := client.SubscribeClipboard(context.Background(), &ClipboardService.SubscribeClipboardRequest{})

	if err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}

	go func() {
		for {
			in, err := stream.Recv()

			if err != nil {
				log.Printf("Error receiving: %v", err)
			}

			log.Printf("Received: %s", in)
			for _, ignoreDeviceId := range ignoreDeviceIdList {
				if in.Items[0].DeviceId == ignoreDeviceId {
					return
				}
			}
			// stdin|pbcopy
			cmd := exec.Command("pbcopy")
			stdin, err := cmd.StdinPipe()
			if err != nil {
				log.Fatal(err)
			}
			go func() {
				defer func(stdin io.WriteCloser) {
					err := stdin.Close()
					if err != nil {
						return
					}
				}(stdin)
				err := cmd.Start()
				if err != nil {
					return
				}
				switch in.Operation {
				case "create":
					{
						_, err = stdin.Write([]byte(strings.Join([]string{in.Items[0].Content}, "")))
						if err != nil {
							return
						}
					}
				case "update":
					{
						_, err = stdin.Write([]byte(strings.Join([]string{in.Items[0].Content}, "")))
						if err != nil {
							return
						}
					}
				case "delete":
					{
						err := cmd.Wait()
						if err != nil {
							return
						}
						return
					}
				}
			}()
		}
	}()
	select {}
}

func retryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	for i := 0; i < 3; i++ {
		err = invoker(ctx, method, req, reply, cc, opts...)
		if err == nil {
			return
		}
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.Unavailable {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(1 * time.Second):
		}
	}
	return nil
}
