package main

import (
	"github.com/texasroh/junecoin/blockchain"
	"github.com/texasroh/junecoin/cli"
	"github.com/texasroh/junecoin/db"
)

func main() {
	defer db.Close()
	blockchain.Blockchain()
	cli.Start()
}
