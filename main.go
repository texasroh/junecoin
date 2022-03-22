package main

import (
	"github.com/texasroh/junecoin/cli"
	"github.com/texasroh/junecoin/db"
)

func main() {
	defer db.Close()
	db.InitDB()
	cli.Start()
}
