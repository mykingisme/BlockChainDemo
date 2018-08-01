package main

import (
	"fmt"
)

func main() {
	bc := CreateBlockChain()
	bc.AddBlock("比特币交易0")
	bc.AddBlock("比特币交易1")

	for _, d := range bc.blocks {
		fmt.Printf("data = %v\n", string(d.Data))
		fmt.Printf("PrevBlockHash = %x\n", (d.PrevBlockHash))
		fmt.Printf("Hash = %x\n", (d.Hash))

	}

}
