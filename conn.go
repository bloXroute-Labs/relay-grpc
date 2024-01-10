package relay_grpc

import (
	"context"
	"fmt"
	"github.com/bloXroute-Labs/relay-grpc/regionalrelay"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func NewRelayConnection(host, authToken string) (regionalrelay.RegionalRelayClient, error) {

	// Check initial connection for approval
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	if conn == nil {
		newConn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println("failed to connect to grpc service with error", "error", err)
			return nil, err
		}

		conn = newConn
	}
	defer conn.Close()

	return regionalrelay.NewRegionalRelayClient(conn), nil
}

func NewConnection(host, authToken string) (chan *regionalrelay.SubmitBlockRequest, error) {

	// Check initial connection for approval
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	bodyChan := make(chan *regionalrelay.SubmitBlockRequest, 100)
	go ConnectToGRPCService(host, authToken, &bodyChan, conn)

	return bodyChan, err
}

func ConnectToGRPCService(host, authToken string, bodyChan *chan *regionalrelay.SubmitBlockRequest, conn *grpc.ClientConn) {
	if conn == nil {
		newConn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	client := regionalrelay.NewRegionalRelayClient(conn)
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
