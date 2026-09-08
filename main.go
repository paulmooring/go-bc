package main

import (
	"fmt"

	"github.com/paulmooring/go-bc/chain"
)

type block struct {
	timestamp	float
	data 		string
	hash 		string
	prev_hash 	string
	nonce		uint
}

func newBlock(data string, prev_hash string) *block {
	blk := block{data: data, prev_hash: prev_hash}
	timestamp := time.time()
	hash = mine()
	nonce = 0
}

func main() {
	blockchain := chain.New()

	fmt.Printf("Created blockchain: %+v\n", blockchain)
}
