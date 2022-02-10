package main

import (
	"github.com/texasroh/junecoin/blockchain"
	"github.com/texasroh/junecoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
