// Copyright 2026 TechDivision GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package style provides terminal output styling (ANSI colors) shared
// across the CLI, so every command renders the same way.
package style

import (
	"io"
	"os"
)

// ANSI codes — same values as the Python callback plugin, so Go-layer
// output stays visually consistent with Ansible task output.
const (
	ansiRed    = "\033[1;31m"
	ansiBlue   = "\033[1;34m"
	ansiGreen  = "\033[0;32m"
	ansiYellow = "\033[0;33m"
	ansiBold   = "\033[1m"
	ansiReset  = "\033[0;0m"
)

// IsTerminal reports whether w is connected to a TTY. When output is piped
// or redirected we skip all ANSI codes, matching fatih/color behavior
// without adding a dependency.
func IsTerminal(w io.Writer) bool {
	if f, ok := w.(*os.File); ok {
		fi, err := f.Stat()
		if err != nil {
			return false
		}
		return (fi.Mode() & os.ModeCharDevice) != 0
	}
	return false
}

// Blue styles s for terminal output, or returns it unchanged when w is not a TTY.
func Blue(w io.Writer, s string) string {
	if IsTerminal(w) {
		return ansiBlue + ansiBold + s + ansiReset
	}
	return s
}

// Green styles s for terminal output, or returns it unchanged when w is not a TTY.
func Green(w io.Writer, s string) string {
	if IsTerminal(w) {
		return ansiGreen + ansiBold + s + ansiReset
	}
	return s
}

// Red styles s for terminal output, or returns it unchanged when w is not a TTY.
func Red(w io.Writer, s string) string {
	if IsTerminal(w) {
		return ansiRed + ansiBold + s + ansiReset
	}
	return s
}

func Yellow(w io.Writer, s string) string {
	if IsTerminal(w) {
		return ansiYellow + ansiBold + s + ansiReset
	}
	return s
}

// Info styles s like Blue, prefixed with an info glyph.
func Info(w io.Writer, s string) string {
	return Blue(w, "ℹ "+s)
}
