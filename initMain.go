package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/spacemeshos/go-spacemesh/common/types"
	_ "github.com/spacemeshos/go-spacemesh/common/util"
)

const testEcho = "test-ping-smtest"

var (
	//Options
	flagInfo = flag.Bool("info", false, "(option) info on address")
	flagEcho = flag.Bool("echo", false, "(option) echo")

	//Params
	flagAddress = flag.String("address", "", "(param) address")
	flagRPCURL  = flag.String("rpcurl", "", "(param) rpc url")

	//Debug
	flagTest = flag.Bool("test", false, "(option) test")
)

var rpcURL string

func test() {
}

func main() {
	flag.Parse()

	if len(*flagRPCURL) == 0 {
		fmt.Println("ERROR: -rpcurl is mandatory")
		os.Exit(1)
	}

	rpcURL = *flagRPCURL

	if *flagEcho {
		v, err := echo(testEcho)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		if v == testEcho {
			fmt.Println("ECHO: OK")
		} else {
			fmt.Println("ERROR: echo mismatch")
		}
	}
}
