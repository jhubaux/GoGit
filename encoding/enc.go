package encoding

import (
	"bytes"
	"github.com/dedis/crypto/abstract"
)

// Packet example
type Packet struct {
	Id     int
	Public abstract.Point
	Secret abstract.Secret
}

func (p Packet) Equal(p2 Packet) bool {
	return p.Id == p2.Id && p.Public.Equal(p2.Public) && p.Secret.Equal(p2.Secret)
}

func ReadPacket(buf []byte) Packet {
	reader := bytes.NewBuffer(buf)
	pack := Packet{}
	testSuite.Read(reader, &pack)
	return pack
}
