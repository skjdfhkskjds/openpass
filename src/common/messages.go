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

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Report reports a message in white
func Report(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// ReportGreen reports a message in green
func ReportGreen(format string, a ...interface{}) {
	color.Green(fmt.Sprintf(format, a...))
}

// ReportYellow reports a message in yellow
func ReportYellow(format string, a ...interface{}) {
	color.Yellow(fmt.Sprintf(format, a...))
}

// PanicRed reports a message in red and exits the program
func PanicRed(format string, a ...interface{}) {
	color.Red(fmt.Sprintf(format, a...))
	os.Exit(1)
}
