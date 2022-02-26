package main

import (
	"fmt"
	"src/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second")
	chain.AddBlock("Third")
	chain.AddBlock("Fourth")

	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
	}
}
