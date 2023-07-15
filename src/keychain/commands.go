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
	"encoding/json"
	"fmt"
	"os"

	. "openpass/common"
)

func (cdc *Codec) Set(domain, username string) string {
	_, found := FindFromJSON(cdc.PasswordsFilePath, domain, username)
	if !found {
		cdc.Delete(domain)
	}

	cdc.SetSalt(cdc.salt)
	password := cdc.generatePassword()
	if err := cdc.setPassword(domain, username, password); err != nil {
		PanicRed(err.Error())
	}

	return password
}

func (cdc *Codec) Get(domain, username string) (string, string) {
	password, found := FindFromJSON(cdc.PasswordsFilePath, domain, username)
	if !found {
		// TODO: so bad
		return "Password for " + domain + " not found", ""
	}

	cdc.SetSalt(password.Salt)
	pass, err := cdc.decrypt(password.Password)
	if err != nil {
		PanicRed(err.Error())
	}

	return username, pass
}

func (cdc *Codec) Update(domain, username, password string) string {
	cdc.Delete(domain)
	cdc.SetSalt(cdc.salt)
	if err := cdc.setPassword(domain, username, password); err != nil {
		PanicRed(err.Error())
	}

	return password
}

func (cdc *Codec) Copy(domain1, username1, domain2, username2 string) string {
	_, password := cdc.Get(domain1, username1)
	return cdc.Update(domain2, username2, password)
}

func (cdc *Codec) Delete(domain string) {
	passwords := ReadJSON(cdc.PasswordsFilePath)

	for i, password := range passwords {
		if password.Domain == domain {
			passwords = Remove(passwords, i)
			break
		}
	}

	jsonPasswords, err := json.Marshal(passwords)
	if err != nil {
		PanicRed(err.Error())
	}
	err = os.WriteFile(cdc.PasswordsFilePath, jsonPasswords, 0644)
	if err != nil {
		PanicRed(err.Error())
	}

	// reset the salt
	cdc.salt = nil
}

func (cdc *Codec) List() string {
	passwords := ReadJSON(cdc.PasswordsFilePath)
	output := ""
	for _, password := range passwords {
		pass, err := cdc.decrypt(password.Password)
		if err != nil {
			PanicRed(err.Error())
		}
		output += "-------------------\n"
		output += fmt.Sprintf("domain: %s\n", password.Domain)
		output += fmt.Sprintf("username: %s\n", password.Username)
		output += fmt.Sprintf("password: %s\n", pass)
		output += "-------------------\n\n"
	}
	return output
}
