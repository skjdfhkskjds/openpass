// SPDX-License-Identifier: AGPL-3.0-only
// Copyright (C) 2023 skjdfhkskjds

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package keychain

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"openpass/types"

	"golang.org/x/crypto/pbkdf2"
)

const (
	iterations = 100000
)

// Codec is the master key used to encrypt and decrypt passwords
type Codec struct {
	key  []byte
	salt []byte

	password string
	*types.Config
}

// GenerateKey generates a new key using PBKDF2
func GenerateCodec(password string, salt []byte, config *types.Config) *Codec {
	// Derive the key using PBKDF2
	key := pbkdf2.Key([]byte(password), salt, iterations, config.KeyLength, sha256.New)

	// Append the salt to the key
	key = append(key, salt...)
	return &Codec{key, salt, password, config}
}

// VerifyPassword verifies that a input matches the encoding password
func (cdc *Codec) VerifyPassword(password string) bool {
	return password == cdc.password
}

// encrypt encrypts the password using the key
func (cdc *Codec) encrypt(password string) (string, error) {
	block, err := aes.NewCipher(cdc.key[:cdc.KeyLength])
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
func (cdc *Codec) decrypt(encodedData string) (string, error) {
	// Decode the base64 encoded data
	encryptedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(cdc.key[:cdc.KeyLength])
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

func (cdc *Codec) SetSalt(salt []byte) {
	if salt == nil {
		salt = types.GenerateSalt(cdc.SaltLength)
		cdc.salt = salt
		cdc.key = append(cdc.key, salt...)
	}
	cdc.salt = salt
}
