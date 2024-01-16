package relay_grpc

import (
	context "context"
	"fmt"
  "time"

	grpc "google.golang.org/grpc"
  "google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func NewRelayConnection(host, authToken string) (RelayClient, error) {
  ka := keepalive.ClientParameters{
      Time:                30 * time.Second,  // send keepalive every 30 seconds
      Timeout:             10 * time.Second,  // wait 10 seconds for ping back
      PermitWithoutStream: true,              // send pings even without active streams
  }

	// Check initial connection for approval
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(ka))
	if err != nil {
		return nil, err
	}
	if conn == nil {
		newConn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(ka))
		if err != nil {
			fmt.Println("failed to connect to grpc service with error", "error", err)
			return nil, err
		}

		conn = newConn
	}
	defer conn.Close()

	return NewRelayClient(conn), nil
}

func NewConnection(host, authToken string) (chan *SubmitBlockRequest, error) {
  ka := keepalive.ClientParameters{
      Time:                30 * time.Second,  // send keepalive every 30 seconds
      Timeout:             10 * time.Second,  // wait 10 seconds for ping back
      PermitWithoutStream: true,              // send pings even without active streams
  }

	// Check initial connection for approval
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(ka))
	if err != nil {
		return nil, err
	}
	bodyChan := make(chan *SubmitBlockRequest, 100)
	go ConnectToGRPCService(host, authToken, &bodyChan, conn)

	return bodyChan, err
}

func ConnectToGRPCService(host, authToken string, bodyChan *chan *SubmitBlockRequest, conn *grpc.ClientConn) {
  ka := keepalive.ClientParameters{
      Time:                30 * time.Second,  // send keepalive every 30 seconds
      Timeout:             10 * time.Second,  // wait 10 seconds for ping back
      PermitWithoutStream: true,              // send pings even without active streams
  }

	if conn == nil {
		newConn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(ka))
		if err != nil {
			fmt.Println("failed to connect to grpc service with error", "error", err)
			return
		}

		conn = newConn
	}
	defer conn.Close()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("grpc panicked but recovered!", "error", r)
			go ConnectToGRPCService(host, authToken, bodyChan, nil)
		} else {
			fmt.Println("grpc closed without panic, restarting to be safe")
			go ConnectToGRPCService(host, authToken, bodyChan, nil)
		}
	}()

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", authToken)

	client := NewRelayClient(conn)
	for {
		body := <-*bodyChan
		go func() {
			_, err := client.SubmitBlock(ctx, body)
			if err != nil {
				fmt.Println("failed to submit block over grpc with error", "error", err)
				return
			}
			fmt.Println("successfully submitted block using grpc")
		}()

	}
}
