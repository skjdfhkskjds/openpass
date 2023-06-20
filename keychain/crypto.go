package keychain

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

const (
	iterations = 100000
	keyLength  = 32 // 32 bytes = 256 bits (AES-256)
	saltLength = 16 // 16 bytes = 128 bits
)

// Codec is the master key used to encrypt and decrypt passwords
type Codec struct {
	key  []byte
	salt []byte
}

// encrypt encrypts the password using the key
func (cdc Codec) encrypt(password string) (string, error) {
	block, err := aes.NewCipher(cdc.key[:keyLength])
	if err != nil {
		return "", err
	}

	// Generate a random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Encrypt the plaintext
	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(password))
	stream.XORKeyStream(ciphertext, []byte(password))

	// Concatenate IV and ciphertext
	encryptedData := append(iv, ciphertext...)

	// Encode the encrypted data to base64 for storage
	encodedData := base64.StdEncoding.EncodeToString(encryptedData)

	return encodedData, nil
}

// decrypt decrypts the base64 encoded ciphertext using the key
func (cdc Codec) decrypt(encodedData string) (string, error) {
	// Decode the base64 encoded data
	encryptedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(cdc.key[:keyLength])
	if err != nil {
		return "", err
	}

	// Extract the IV from the encrypted data
	iv := encryptedData[:aes.BlockSize]
	ciphertext := encryptedData[aes.BlockSize:]

	// Decrypt the ciphertext
	stream := cipher.NewCFBDecrypter(block, iv)
	password := make([]byte, len(ciphertext))
	stream.XORKeyStream(password, ciphertext)

	return string(password), nil
}

// GenerateKey generates a new key using PBKDF2
func GenerateCodec(password string, salt []byte) *Codec {
	// Derive the key using PBKDF2
	key := pbkdf2.Key([]byte(password), salt, iterations, keyLength, sha256.New)

	// Append the salt to the key
	key = append(key, salt...)
	return &Codec{key, salt}
}

// GenerateSalt returns a random salt of saltLength bytes
func GenerateSalt() []byte {
	salt := make([]byte, saltLength)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		log.Fatal(err)
	}
	return salt
}

func (cdc *Codec) SetSalt(salt []byte) {
	if salt == nil {
		salt = GenerateSalt()
		cdc.key = append(cdc.key, salt...)
	}
	cdc.salt = salt
}
