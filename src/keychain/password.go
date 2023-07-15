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
	"math/rand"
	"time"
)

const (
	digits = "0123456789"
	alpha  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz"
)

type Password struct {
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     []byte `json:"salt"`
}

// generatePassword generates a password with at least one
// digit and one special character
func (cdc *Codec) generatePassword() string {
	allChars := alpha + digits + cdc.SpecialCharaters

	rand.Seed(time.Now().UnixNano())
	passBytes := make([]byte, cdc.PasswordLength)
	passBytes[0] = digits[rand.Intn(len(digits))]
	passBytes[1] = cdc.SpecialCharaters[rand.Intn(len(cdc.SpecialCharaters))]
	for i := 2; i < cdc.PasswordLength; i++ {
		passBytes[i] = allChars[rand.Intn(len(allChars))]
	}
	rand.Shuffle(len(passBytes), func(i, j int) {
		passBytes[i], passBytes[j] = passBytes[j], passBytes[i]
	})
	password := string(passBytes)
	return password
}

func (cdc *Codec) setPassword(domain, username, password string) error {
	encryptedPassword, err := cdc.encrypt(password)
	if err != nil {
		return err
	}

	result := Password{
		Domain:   domain,
		Username: username,
		Password: encryptedPassword,
		Salt:     cdc.salt,
	}

	if err := AppendJSON(cdc.PasswordsFilePath, result); err != nil {
		return err
	}

	return nil
}

func Remove(slice []Password, i int) []Password {
	slice[i] = slice[len(slice)-1]   // Copy last element to index i.
	slice[len(slice)-1] = Password{} // Erase last element (write zero value).
	return slice[:len(slice)-1]
}
