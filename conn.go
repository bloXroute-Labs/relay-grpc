package relay_grpc

import (
	context "context"
	"fmt"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func NewConnection(host, authToken string) (chan *SubmitBlockRequest, error) {

	// Check initial connection for approval
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	bodyChan := make(chan *SubmitBlockRequest, 100)
	go ConnectToGRPCService(host, authToken, &bodyChan, conn)

	return bodyChan, err
}

func ConnectToGRPCService(host, authToken string, bodyChan *chan *SubmitBlockRequest, conn *grpc.ClientConn) {
	if conn == nil {
		newConn, err := grpc.Dial(host, grpc.WithInsecure())
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
