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

	//Params
	flagAddress = flag.String("address", "", "(param) address")

	//Debug
	flagTest = flag.Bool("test", false, "(option) test")
)

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
