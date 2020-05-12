package main

import (
	"context"
	"time"

	sm "github.com/inc4/sm-test/spacemesh"
	"google.golang.org/grpc"
)

func echo(s string) (value string, err error) {
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
