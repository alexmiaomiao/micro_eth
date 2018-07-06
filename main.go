package main

import (
	"fmt"

	"github.com/alexmiaomiao/micro_eth/core"
	"github.com/alexmiaomiao/micro_eth/miner"
)

func main() {
	bc := core.NewBlockchain()

	me := miner.New("Miner lzt win 50 BTC", bc)
	me.Start("jm Send 1 BTC to ld")
	fmt.Println()
	me.Start("ld Send 2 BTC to jm||ld Send 3 BTC to lzt")
	fmt.Println()

	for _, block := range bc.Blocks() {
		fmt.Printf("Number: %d\n", block.Header.Number)
		fmt.Printf("Parenthash: %x\n", block.Header.ParentHash)
		fmt.Printf("Hash: %x\n", block.Header.SelfHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println()
	}
}
