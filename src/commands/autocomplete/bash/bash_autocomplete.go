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

package bash

import (
	"os"
	"os/exec"
	"path/filepath"

	"openpass/common"

	"github.com/spf13/cobra"
)

const (
	bashConfigPath = ".bashrc"
)

type Shell struct {
	FilePath string
}

func (sh Shell) GenerateAutoCompleteScripts(cmd *cobra.Command) error {
	if err := cmd.GenZshCompletionFile(sh.FilePath); err != nil {
		return err
	}
	return sh.sourceScriptForBash()
}

// sourceScriptForBash sources the completion script for Bash
// and refreshes the Bash shell
func (sh Shell) sourceScriptForBash() error {
	bashrc, err := sh.getBashConfig()
	if err != nil {
		return err
	}

	// Write the completion script to a file in the Zsh config directory
	data := []byte("source " + sh.FilePath + "\n")
	if err := common.WriteFile(bashrc, data); err != nil {
		return err
	}

	return exec.Command("source " + bashrc).Run()
}

// getBashConfig returns the path to the Bash config file
func (sh Shell) getBashConfig() (string, error) {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, bashConfigPath), nil
}
