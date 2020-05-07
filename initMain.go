package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/spacemeshos/go-spacemesh/common/types"
	_ "github.com/spacemeshos/go-spacemesh/common/util"
)

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

	if len(*flagRPCURL) == 0 {
		fmt.Println("ERROR: -rpcurl is mandatory")
		os.Exit(1)
	}

	rpcURL = *flagRPCURL

	if *flagEcho {
		v, err := echo("test-ping")
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		fmt.Printf("ECHO: %s\n", v)
	}

	if *flagInfo {
		if len(*flagAddress) == 0 {
			fmt.Println("ERROR: -address is mandatory for info")
			os.Exit(1)
		}

		rpcURL = *flagRPCURL

		addressInfo(*flagAddress)
	}
}
