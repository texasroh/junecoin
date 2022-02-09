package db

import (
	"github.com/boltdb/bolt"
	"github.com/texasroh/junecoin/utils"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		db, err := bolt.Open("blockchain.db", 0600, nil)
		utils.HandleErr(err)

		db.dbPointer
	}
	return db
}
