package main

import (
	"github.com/texasroh/junecoin/blockchain"
)

func main() {
	blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")
}
