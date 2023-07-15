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

package commands

import (
	"syscall"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"

	"openpass/commands/autocomplete"
	"openpass/keychain"
	"openpass/types"

	. "openpass/common"
)

var (
	domain   string
	username string

	cdc *keychain.Codec
)

var (
	rootCmd = &cobra.Command{
		Use:   "keys",
		Short: "keys is a CLI application which manages your passwords",
		Run: func(cmd *cobra.Command, args []string) {
			PanicRed("Please specify a valid command. Use --help for more information.")
		},
	}

	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set a password",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				PanicRed(InvalidCommand(SetCommandArgs))
			}

			if cdc.RequirePasswordCheck {
				ReportYellow("Verify your password: ")
				for {
					inputPassword, err := terminal.ReadPassword(int(syscall.Stdin))
					if err != nil {
						PanicRed(err.Error())
					}
					if cdc.VerifyPassword(string(inputPassword)) {
						break
					}
					ReportYellow(MismatchedPasswords)
				}
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			password := cdc.Set(domain, username)
			ReportYellow("Setting password...")
			outputResult(username, password)
		},
	}

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get a password",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				PanicRed(InvalidCommand(GetCommandArgs))
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			ReportYellow("Getting password...")
			username, password := cdc.Get(domain, username)
			outputResult(username, password)
		},
	}

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a password",
		Args:  cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				PanicRed(InvalidCommand(UpdateCommandArgs))
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			newPassword := args[0]
			ReportYellow("Updating password...")
			cdc.Update(domain, username, newPassword)
			outputResult(username, newPassword)
		},
	}

	copyCmd = &cobra.Command{
		Use:   "copy",
		Short: "Copies a password",
		Args:  cobra.ExactArgs(2),
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				PanicRed(InvalidCommand(CopyCommandArgs))
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			destDomain := args[0]
			destUsername := args[1]
			ReportYellow("Copying password from %s to %s\n", domain, destDomain)
			password := cdc.Copy(domain, username, destDomain, destUsername)
			outputResult(destUsername, password)
		},
	}

	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a password",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				PanicRed(InvalidCommand(DeleteCommandArgs))
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			ReportGreen("Deleting password...")
			cdc.Delete(domain)
		},
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all passwords",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				PanicRed(InvalidCommand(ListCommandArgs))
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			ReportGreen("Listing passwords...")
			ReportCyan(cdc.List())
		},
	}
)

// SetupCommands sets up the commands for the CLI
func SetupCommands(userPassword string) *cobra.Command {
	config, err := types.LoadConfig()
	if err != nil {
		PanicRed(err.Error())
	}
	cdc = keychain.GenerateCodec(userPassword, nil, config)

	// Set up pre-run
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if cmd.Name() == rootCmd.Name() || cmd.Name() == listCmd.Name() {
			return
		}

		if domain == "" {
			PanicRed(InvalidCommand(BaseCommandArgs))
		}
		if username == "" {
			PanicRed(InvalidCommand(BaseCommandArgs))
		}

		domain = ReduceDomain(domain)
	}

	// Add flags
	rootCmd.PersistentFlags().StringVar(&domain, "domain", "", "Domain name")
	getCmd.PersistentFlags().StringVar(&username, "user", "", "Username")

	// Add all commands to root
	rootCmd.AddCommand(setCmd, getCmd, updateCmd, copyCmd, deleteCmd, listCmd)

	// Generate shell completion scripts
	if config.GenerateAutoCompleteScripts {
		autocomplete.GenerateFilePath(rootCmd)
	}

	return rootCmd
}

// outputResult outputs the username and password to the user
func outputResult(username, password string) {
	ReportGreen("Username: %s", username)
	if !cdc.HidePassword {
		ReportGreen("Password: %s", password)
	}
	if err := clipboard.WriteAll(password); err != nil {
		PanicRed(err.Error())
	}
	ReportGreen("Password copied to clipboard")
}
