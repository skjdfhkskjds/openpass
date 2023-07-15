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
	"io/ioutil"
	"openpass/common"
	"os"
)

func AppendJSON(jsonFile string, password Password) error {
	passwords := ReadJSON(jsonFile)

	passwords = append(passwords, password)

	jsonPasswords, err := json.MarshalIndent(passwords, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonFile, jsonPasswords, 0644)
	if err != nil {
		return err
	}

	return nil
}

// findFromJSON finds a password from a JSON file
func FindFromJSON(jsonFile, domain, username string) (Password, bool) {
	passwords := ReadJSON(jsonFile)

	for _, password := range passwords {
		if password.Domain == domain && password.Username == username {
			return password, true
		}
	}

	// Read the JSON file into a byte array
	return Password{}, false
}

// readJSON reads a JSON file and returns a slice of Passwords
func ReadJSON(jsonFile string) []Password {
	// Open the JSON file
	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil || len(jsonData) == 0 {
		return []Password{}
	}

	var passwords []Password
	if err := json.Unmarshal(jsonData, &passwords); err != nil {
		common.PanicRed(err.Error())
	}

	return passwords
}
