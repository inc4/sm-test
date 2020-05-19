package main

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/inc4/sm-test/spacemesh"
	"golang.org/x/exp/errors/fmt"
	"google.golang.org/grpc"
)

func testEcho(s string) (value string, err error) {
	conn, err := grpc.Dial(rpcURL, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := spacemesh.NewNodeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	v := spacemesh.SimpleString{Value: s}

	res, err := c.Echo(ctx, &v)
	if err != nil {
		return
	}

	value = res.GetValue()

	return
}

func testStartSync() (err error) {
	conn, err := grpc.Dial(rpcURL, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := spacemesh.NewNodeServiceClient(conn)

	_, err = c.SyncStart(context.Background(), &empty.Empty{})
	if err != nil {
		return
	}

	return
}

func testNodeSync() (err error) {
	err = testStartSync()
	if err != nil {
		return
	}

	conn, err := grpc.Dial(rpcURL, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := spacemesh.NewNodeServiceClient(conn)

	stream, err := c.SyncStatusStream(context.Background(), &empty.Empty{})
	if err != nil {
		return
	}

	fmt.Println("testNodeSync: Connected, listening ...")

	for {
		status, err := stream.Recv()
		if err != nil {
			fmt.Printf("Recv(ERROR): %v\n", err)

			break
		}

		fmt.Printf("Recv(OK): %s\n", status.Status.String())
	}

	return
}

func testMeshLayer() (err error) {
	err = testStartSync()
	if err != nil {
		return
	}

	conn, err := grpc.Dial(rpcURL, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := spacemesh.NewMeshServiceClient(conn)

	stream, err := c.LayerStream(context.Background(), &empty.Empty{})
	if err != nil {
		return
	}

	fmt.Println("testMeshLayer: Connected, listening ...")

	for {
		layer, err := stream.Recv()
		if err != nil {
			fmt.Printf("Recv(ERROR): %v\n", err)

			break
		}

		fmt.Printf("Recv(OK): %d - %s\n", layer.GetNumber(), layer.GetStatus())
	}

	return
}

func testGlobalTx() (err error) {
	err = testStartSync()
	if err != nil {
		return
	}

	conn, err := grpc.Dial(rpcURL, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := spacemesh.NewGlobalStateServiceClient(conn)

	stream, err := c.TransactionReceiptStream(context.Background(), &empty.Empty{})
	if err != nil {
		return
	}

	fmt.Println("testGlobalTx: Connected, listening ...")

	for {
		txReceipt, err := stream.Recv()
		if err != nil {
			fmt.Printf("Recv(ERROR): %v\n", err)

			break
		}

		fmt.Printf("Recv(OK): %d - %s - %s\n",
			txReceipt.GetLayerNumber(),
			hex.EncodeToString(txReceipt.GetId().Id),
			txReceipt.GetResult().String(),
		)
	}

	return
}
