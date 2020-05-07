package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	sm "github.com/inc4/sm-test/spacemesh"
	_ "github.com/spacemeshos/go-spacemesh/common/types"
	_ "github.com/spacemeshos/go-spacemesh/common/util"
	"google.golang.org/grpc"
)

var (
	//Options
	flagInfo = flag.Bool("info", false, "(option) info on address")

	//Params
	flagAddress = flag.String("address", "", "(param) address")

	//Debug
	flagTest = flag.Bool("test", false, "(option) test")
)

func echo(s string) (value string, err error) {
	conn, err := grpc.Dial("localhost:9990", grpc.WithInsecure())
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

func test() {
	v, err := echo("test-ping")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	fmt.Printf("ECHO: %s\n", v)
}

func addressInfo(address string) {

	bal, err := getBalance(address)
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := getNonce(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address, bal, nonce)
}

func main() {
	flag.Parse()

	if *flagTest {
		test()
	}

	if *flagInfo {
		if len(*flagAddress) == 0 {
			fmt.Println("ERROR: -address is mandatory for info")
			os.Exit(1)
		}

		addressInfo(*flagAddress)
	}
}
