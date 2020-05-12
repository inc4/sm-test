package main

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	sm "github.com/inc4/sm-test/spacemesh"
	"golang.org/x/exp/errors/fmt"
	"google.golang.org/grpc"
)

func testEcho(s string) (value string, err error) {
	conn, err := grpc.Dial(rpcURL, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := sm.NewNodeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	v := sm.SimpleString{Value: s}

	res, err := c.Echo(ctx, &v)
	if err != nil {
		return
	}

	value = res.GetValue()

	return
}

func testSync() (value sm.NodeService_SyncStatusStreamClient, err error) {
	conn, err := grpc.Dial(rpcURL, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := sm.NewNodeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	value, err = c.SyncStatusStream(ctx, &empty.Empty{})
	if err != nil {
		return
	}

	fmt.Printf("SyncStatusStream: %v\n", value)

	for {
		nodeStatus, err := value.Recv()
		if err != nil {
			fmt.Printf("Recv(ERROR): %v\n", err)

			break
		}

		fmt.Printf("Recv(OK): %v\n", nodeStatus)
	}

	return
}
