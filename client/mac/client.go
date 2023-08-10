package main

import (
	"context"
	ClipboardService "github.com/1939323749/clipboard_server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os/exec"
	"strings"
	"time"
)

type SETTING struct {
	tryConnectInterval time.Duration
	server             string
	port               string
	checkAliveInterval time.Duration
}

var ignoreDeviceIdList = []string{"macOS_popclip"}

func main() {
	setting := SETTING{
		tryConnectInterval: 5 * time.Second,
		server:             "localhost",
		port:               "50051",
		checkAliveInterval: 5 * time.Second,
	}

	conn, err := grpc.Dial(setting.server+":"+setting.port, grpc.WithTransportCredentials(insecure.NewCredentials()))

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
	alive, err := client.CheckConnectivity(context.Background())

	if err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}

	go func() {
		ticker := time.NewTicker(setting.checkAliveInterval)
		for range ticker.C {
			if stream == nil {
				continue
			}
			err := alive.Send(&ClipboardService.Alive{})
			if err != nil {
				log.Printf("Error sending: %v", err)
				go func() {
					for {
						time.Sleep(setting.tryConnectInterval)
						conn, err = grpc.Dial(setting.server+":"+setting.port, grpc.WithTransportCredentials(insecure.NewCredentials()))
						if err != nil {
							log.Printf("did not connect: %v", err)
						}
						client = ClipboardService.NewClipboardServiceClient(conn)
						stream, err = client.SubscribeClipboard(context.Background(), &ClipboardService.SubscribeClipboardRequest{})
						alive, err = client.CheckConnectivity(context.Background())
						if err != nil {
							log.Printf("Error subscribing: %v", err)
						}
					}
				}()
				continue
			}
			_, err = alive.Recv()
			if err != nil {
				log.Printf("Error receiving: %v", err)
			} else {
				log.Printf("Alive")
			}
		}
	}()

	go func() {
		for {
			if stream == nil {
				continue
			}
			in, err := stream.Recv()
			if err != nil {
				log.Printf("Error receiving: %v", err)
				time.Sleep(1 * time.Second)
				continue
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
