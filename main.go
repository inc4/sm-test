package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/spacemeshos/go-spacemesh/common/types"
	_ "github.com/spacemeshos/go-spacemesh/common/util"
)

const echoText = "test-ping-smtest"

var (
	//Options

	//Params
	flagRPCURL = flag.String("rpcurl", "", "(param) rpc url")

	//Debug
	flagTestEcho          = flag.Bool("test_echo", false, "(option) test echo")
	flagTestNodeSync      = flag.Bool("test_nodesync", false, "(option) node sync stream")
	flagTestMeshLayer     = flag.Bool("test_meshlayer", false, "(option) mesh layer stream")
	flagTestGlobalTx      = flag.Bool("test_globaltx", false, "(option) global transaction receipts")
	flagTestGlobalRewards = flag.Bool("test_globalrewards", false, "(option) global rewards")
)

var rpcURL string

func main() {
	flag.Parse()

	if len(*flagRPCURL) == 0 {
		fmt.Println("ERROR: -rpcurl is mandatory")
		os.Exit(1)
	}

	rpcURL = *flagRPCURL

	if *flagTestEcho {
		v, err := testEcho(echoText)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		if v == echoText {
			fmt.Println("ECHO: OK")
		} else {
			fmt.Println("ERROR: echo mismatch")
		}

		os.Exit(1)
	}

	if *flagTestNodeSync {
		err := testNodeSync()
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		os.Exit(1)
	}

	if *flagTestMeshLayer {
		err := testMeshLayer()
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		os.Exit(1)
	}

	if *flagTestGlobalTx {
		err := testGlobalTx()
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		os.Exit(1)
	}

	if *flagTestGlobalRewards {
		err := testGlobalRewards()
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		os.Exit(1)
	}
}
