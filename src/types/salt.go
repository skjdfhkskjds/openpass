package types

import (
	"crypto/rand"
	"io"
	"log"
)

const (
	saltLength = 16 // 16 bytes = 128 bits
)

// GenerateSalt returns a random salt of saltLength bytes
func GenerateSalt() []byte {
	salt := make([]byte, saltLength)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		log.Fatal(err)
	}
	return salt
}
