package main

import (
	"context"
	"time"

	pb "github.com/inc4/sm-test/pb"
	"google.golang.org/grpc"
)

const rpcURL = "localhost:9091"

func getBalance(address string) (value string, err error) {
	conn, err := grpc.Dial(rpcURL, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	c := pb.NewSpacemeshServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	accountID := pb.AccountId{Address: address}
	//opts := grpc.CallOption{}
	res, err := c.GetBalance(ctx, &accountID)
	if err != nil {
		return
	}
	value = res.GetValue()
	if len(value) > 12 {
		dp := len(value) - 12
		value = value[0:dp] + "." + value[dp:]
	}
	return

}

func getNonce(address string) (value string, err error) {
	conn, err := grpc.Dial(rpcURL, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	c := pb.NewSpacemeshServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	accountID := pb.AccountId{Address: address}
	//opts := grpc.CallOption{}
	res, err := c.GetNonce(ctx, &accountID)
	if err != nil {
		return
	}
	value = res.GetValue()
	return
}
