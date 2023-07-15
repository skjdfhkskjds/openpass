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

package common

import "fmt"

const baseErrorMessage = "Invalid format\nexpected:"

// Command Arg Messages
const (
	MismatchedPasswords = "Passwords do not match. Try again..."
	UnsupportedShell    = "Unsupported shell, unable to generate autocomplete scripts."

	BaseCommandArgs   = "command <arg> [ARGS]"
	SetCommandArgs    = "set --domain <domain> --user <username>"
	GetCommandArgs    = "get --domain <domain> --user <username>"
	UpdateCommandArgs = "update --domain <domain> --user <username> <password>"
	CopyCommandArgs   = "copy --domain <src-domain> --user <src-username> <dest-domain> <dest-username>"
	DeleteCommandArgs = "delete --domain <domain> --user <username>"
	ListCommandArgs   = "list"
)

// Config Messages
const (
	PasswordLengthMessage = "password length must be greater than 0"
	SaltLengthMessage     = "salt length must be greater than 0"

	MissingConfigPath = "missing config file path, generating..."
)

func InvalidCommand(message string) string {
	return fmt.Sprintf("%s %s", baseErrorMessage, message)
}
