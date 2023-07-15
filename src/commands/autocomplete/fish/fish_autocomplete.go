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

package fish

import (
	"openpass/common"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	fishConfigPath = ".config/fish/config.fish"
)

type Shell struct {
	FilePath string
}

func (sh Shell) GenerateAutoCompleteScripts(cmd *cobra.Command) error {
	if err := cmd.GenZshCompletionFile(sh.FilePath); err != nil {
		return err
	}
	return sh.sourceScriptForFish()
}

// sourceScriptForFish sources the completion script for Fish
// and refreshes the Fish shell
func (sh Shell) sourceScriptForFish() error {
	fishConfig, err := sh.getFishConfig()
	if err != nil {
		return err
	}

	// Write the completion script to a file in the Zsh config directory
	data := []byte("source " + sh.FilePath + "\n")
	if err := common.WriteFile(fishConfig, data); err != nil {
		return err
	}

	return exec.Command("source " + fishConfig).Run()
}

// getFishConfig returns the path to the Fish config file
func (sh Shell) getFishConfig() (string, error) {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, fishConfigPath), nil
}
