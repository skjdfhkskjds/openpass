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

package autocomplete

import (
	"os"
	"strings"

	"openpass/commands/autocomplete/bash"
	"openpass/commands/autocomplete/fish"
	"openpass/commands/autocomplete/powershell"
	"openpass/commands/autocomplete/zsh"
	. "openpass/common"

	"github.com/spf13/cobra"
)

const (
	bashCompletionFile       = "../keys_autocomplete.sh"
	zshCompletionFile        = "../keys_autocomplete.sh"
	fishCompletionFile       = "../keys_autocomplete.fish"
	powerShellCompletionFile = "../keys_autocomplete.ps1"
)

type Shell interface {
	GenerateAutoCompleteScripts(cmd *cobra.Command) error
}

func GenerateFilePath(cmd *cobra.Command) error {
	sh := newShell()
	if err := sh.GenerateAutoCompleteScripts(cmd); err != nil {
		return err
	}
	ReportGreen(
		"Successfully generated autocomplete scripts\n for %s",
		os.Getenv("SHELL"),
	)
	return nil
}

func newShell() Shell {
	shell := strings.ToLower(os.Getenv("SHELL"))
	switch {
	case strings.Contains(shell, "zsh"):
		return zsh.Shell{
			FilePath: GetPathOfCaller(zshCompletionFile),
		}

	case strings.Contains(shell, "bash"):
		return bash.Shell{
			FilePath: GetPathOfCaller(bashCompletionFile),
		}

	case strings.Contains(shell, "fish"):
		return fish.Shell{
			FilePath: GetPathOfCaller(fishCompletionFile),
		}

	case strings.Contains(shell, "powershell"):
		return powershell.Shell{
			FilePath: GetPathOfCaller(powerShellCompletionFile),
		}

	default:
		ReportYellow(UnsupportedShell)
		return nil
	}
}
