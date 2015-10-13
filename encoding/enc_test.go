package encoding

import (
	"bytes"
	"github.com/dedis/crypto/config"
	"github.com/dedis/crypto/edwards"
	"github.com/dedis/crypto/random"
	"testing"
)

var testSuite = edwards.NewAES128SHA256Ed25519(true)

// You must read a packet given a byte stream and return the decoded struct
func TestRead(t *testing.T) {

	cf := config.KeyPair{}
	cf.Gen(testSuite, random.Stream)

	pack := Packet{
		Id:     1,
		Public: cf.Public,
		Secret: cf.Secret,
	}

	var buf bytes.Buffer
	err := testSuite.Write(&buf, &pack)
	if err != nil {
		t.Error("oups. Something wrong on my side... ")
	}

	// Func to implement
	readPacket := ReadPacket(buf.Bytes())

	if !readPacket.Equal(pack) {
		t.Error("Packet is not equal to the marshalled one")
	}
}
