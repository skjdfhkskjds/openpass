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

package powershell

import (
	"openpass/common"
	"os/exec"

	"github.com/spf13/cobra"
)

type Shell struct {
	FilePath string
}

func (sh Shell) GenerateAutoCompleteScripts(cmd *cobra.Command) error {
	if err := cmd.GenZshCompletionFile(sh.FilePath); err != nil {
		return err
	}
	return sh.sourceScriptForPowerShell()
}

// sourceScriptForPowerShell sources the completion script for PowerShell
// and refreshes the PowerShell shell
func (sh Shell) sourceScriptForPowerShell() error {
	psConfig, err := getPowershellConfig()
	if err != nil {
		return err
	}

	data := []byte(". " + sh.FilePath + "\n")
	if err := common.WriteFile(psConfig, data); err != nil {
		return err
	}

	return exec.Command(". " + psConfig).Run()
}

// getPowershellConfig returns the path to the PowerShell profile directory
func getPowershellConfig() (string, error) {
	cmd := exec.Command("write-output", "$PROFILE")
	psConfig, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(psConfig), nil
}
