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
	flagTestEcho = flag.Bool("test_echo", false, "(option) test echo")
	flagTestSync = flag.Bool("test_sync", false, "(option) sync")
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
	}

	if *flagTestSync {
		_, err := testSync()
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
