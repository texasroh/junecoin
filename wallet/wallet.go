package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/texasroh/junecoin/utils"
)

const (
	signature     string = "3dfd7eace0e02fc074da55c7a8c3ba3d7b821e10b5825af01f09eb57cfbb07cd8928a55d26e18a552f056f7d4df32b5b18b3e90fe30c75dbacb58d97ce9e7ba5"
	privateKey    string = "3077020101042086bb44cd7c927cecd4b4304bb8dd22803769c2da9dd04c7de2d92288cfaa63caa00a06082a8648ce3d030107a1440342000471812e44ceffe112fb6988866307e25b9541e2ce59fda0f886c3dc84d8b675c21377337e3c8f00fd441a1c4b642a3231f8c26e496bb7146c4a8a5b6005db07bb"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {
	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)
	_, err = x509.ParseECPrivateKey(privBytes)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	utils.HandleErr(err)

	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)
	fmt.Println(bigR)
	fmt.Println(bigS)
}
