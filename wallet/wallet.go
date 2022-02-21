package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/texasroh/junecoin/utils"
)

const (
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("%x", signature)
	utils.HandleErr(err)

}
