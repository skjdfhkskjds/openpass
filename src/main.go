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

package main

import (
	"openpass/commands"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	. "openpass/common"
)

// TODO: add tests, dynamically update config file
func main() {
	ReportGreen("Enter your password: ")
	inputPassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		PanicRed(err.Error())
	}

	commands := commands.SetupCommands(string(inputPassword))
	if err := commands.Execute(); err != nil {
		PanicRed(err.Error())
	}
}
