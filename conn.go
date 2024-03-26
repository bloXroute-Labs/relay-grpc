package relay_grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
)

// GRPC dial options
const (
	windowSize = 128 * 1024 * 10
	bufferSize = 0
)

var DefaultKeepaliveParams = keepalive.ClientParameters{
	Time:                30 * time.Second, // send pings every 30 seconds if there is no activity
	Timeout:             20 * time.Second, // wait 20 seconds for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func NewRelayConnection(host string) (RelayClient, error) {
	dialOptions := []grpc.DialOption{
		grpc.WithInitialWindowSize(windowSize),
		grpc.WithInitialConnWindowSize(windowSize),
		//grpc.WithWriteBufferSize(bufferSize),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(DefaultKeepaliveParams),
	}

	// Check initial connection for approval
	conn, err := grpc.Dial(host, dialOptions...)
	if err != nil {
		return nil, err
	}
	if conn == nil {
		newConn, err := grpc.Dial(host, dialOptions...)
		if err != nil {
			fmt.Println("failed to connect to grpc service with error", "error", err)
			return nil, err
		}

		conn = newConn
	}

	return NewRelayClient(conn), nil
}

func NewConnection(host, authToken string, useGzipCompression bool) (chan *SubmitBlockRequest, error) {
	dialOptions := []grpc.DialOption{
		grpc.WithInitialWindowSize(windowSize),
		grpc.WithInitialConnWindowSize(windowSize),
		//grpc.WithWriteBufferSize(bufferSize),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(DefaultKeepaliveParams),
	}

	if useGzipCompression {
		dialOptions = append(dialOptions, grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	}

	// Check initial connection for approval
	conn, err := grpc.Dial(host, dialOptions...)
	if err != nil {
		return nil, err
	}
	bodyChan := make(chan *SubmitBlockRequest, 100)
	go ConnectToGRPCService(host, authToken, &bodyChan, conn, dialOptions)

	return bodyChan, err
}

func ConnectToGRPCService(host, authToken string, bodyChan *chan *SubmitBlockRequest, conn *grpc.ClientConn, dialOptions []grpc.DialOption) {
	if conn == nil {
		newConn, err := grpc.Dial(host, dialOptions...)
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
			go ConnectToGRPCService(host, authToken, bodyChan, nil, dialOptions)
		} else {
			fmt.Println("grpc closed without panic, restarting to be safe")
			go ConnectToGRPCService(host, authToken, bodyChan, nil, dialOptions)
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
