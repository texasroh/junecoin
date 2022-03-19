package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"io/fs"
	"reflect"
	"testing"
)

const (
	testKey     string = "307702010104202287ed39824099882e8d9a2ba89703d0e93cae9bfcfca851f60b0f7eab27675ba00a06082a8648ce3d030107a14403420004202d2e83b96cd84313ea2c2796228ae38e8e9fcaf360034839134d48d00982e9e524e72a464af3e25d4307dd20aa5e2958db2c4881b78077546b7ef3e3adc0b7"
	testPayload string = "000000f23c02e44c8bf911dae5f090c8121bb6ff5debc75be009d6a81b77a8c5"
	testSig     string = "3f764765513655aeeebe12706cb8e5ae21f8c8e5c63e04f65059d003dc289c539bdc0faf7d5e18db6c75936103d7dc22e7969b431a1066032a63f4d6ddc0bb54"
)

type fakeLayer struct {
	fakeHasWalletFile func() bool
}

func (f fakeLayer) hasWalletFile() bool {
	return f.fakeHasWalletFile()
}

func (fakeLayer) writeFile(name string, data []byte, perm fs.FileMode) error {
	return nil
}

func (fakeLayer) readFile(name string) ([]byte, error) {
	// return utils.ToBytes(makeTestWallet().privateKey), nil
	return x509.MarshalECPrivateKey(makeTestWallet().privateKey)
}

func TestWallet(t *testing.T) {
	t.Run("New Wallet is created", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return false },
		}
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
	t.Run("Wallet is restored", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return true },
		}
		w = nil
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("Restored Wallet should return a restored wallet instance")
		}
	})
}

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
	type test struct {
		input string
		ok    bool
	}
	tests := []test{
		{testPayload, true},
		{"400000f23c02e44c8bf911dae5f090c8121bb6ff5debc75be009d6a81b77a8c5", false},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil {
		t.Error("restoreBigInts should return error when payload is not hex")
	}
}
