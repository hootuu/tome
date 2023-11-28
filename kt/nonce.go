package kt

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

type Nonce uint64

func (n Nonce) S() string {
	return fmt.Sprintf("%d", n)
}

func NewNonce() Nonce {
	var randomNumber uint64
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return 0
	}
	randomNumber = binary.BigEndian.Uint64(randomBytes)
	return Nonce(randomNumber)
}
