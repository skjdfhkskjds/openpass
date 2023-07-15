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

package types

import (
	"encoding/json"
	"os"

	. "openpass/common"
)

const (
	configFilePath = "../../config.json"
)

// Config contains the config file data
type Config struct {
	RequirePasswordCheck        bool `json:"requirePasswordCheck"`
	GenerateAutoCompleteScripts bool `json:"generateAutoCompleteScripts"`
	HidePassword                bool `json:"hidePassword"`

	PasswordLength int `json:"passwordLength"`
	KeyLength      int `json:"keyLength"`
	SaltLength     int `json:"saltLength"`

	PasswordsFilePath string `json:"passwordsFilePath"`
	SpecialCharaters  string `json:"specialCharacters"`
}

// LoadConfig loads the config file from the specified path
func LoadConfig() (*Config, error) {
	if file, err := os.ReadFile(GetPathOfCaller(configFilePath)); err == nil {
		config, err := readConfig(file)
		if err != nil {
			return &Config{}, err
		}
		return config, nil
	}

	ReportGreen(MissingConfigPath)
	return saveDefaultConfig(), nil
}

// saveDefaultConfig generates a default config file and saves it
// to the specified path
func saveDefaultConfig() *Config {
	config := newDefaultConfig()

	configData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		PanicRed(err.Error())
	}

	if err := os.WriteFile(GetPathOfCaller(configFilePath), configData, 0644); err != nil {
		PanicRed(err.Error())
	}

	return config
}

// newDefaultConfig returns a default config
func newDefaultConfig() *Config {
	return &Config{
		RequirePasswordCheck:        true,
		GenerateAutoCompleteScripts: false,
		HidePassword:                true,

		PasswordLength: 32,
		SaltLength:     16,
		KeyLength:      32, // 32 bytes = 256 bits (AES-256)

		PasswordsFilePath: "passwords.json",
		SpecialCharaters:  "!@#$%^&*()_+{}[]:<>?/",
	}
}

// readConfig reads a config file and returns a Config struct
func readConfig(configData []byte) (*Config, error) {
	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		return &Config{}, err
	}

	return &config, nil
}
