package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"testing"
)

const (
	testKey     string = "307702010104202287ed39824099882e8d9a2ba89703d0e93cae9bfcfca851f60b0f7eab27675ba00a06082a8648ce3d030107a14403420004202d2e83b96cd84313ea2c2796228ae38e8e9fcaf360034839134d48d00982e9e524e72a464af3e25d4307dd20aa5e2958db2c4881b78077546b7ef3e3adc0b7"
	testPayload string = "000000f23c02e44c8bf911dae5f090c8121bb6ff5debc75be009d6a81b77a8c5"
	testSig     string = "3f764765513655aeeebe12706cb8e5ae21f8c8e5c63e04f65059d003dc289c539bdc0faf7d5e18db6c75936103d7dc22e7969b431a1066032a63f4d6ddc0bb54"
)

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() sould return a hex encoded string, git %s", s)
	}
}

func TestVerify(t *testing.T) {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	b, _ := x509.MarshalECPrivateKey(privKey)
	t.Logf("%x", b)

}
