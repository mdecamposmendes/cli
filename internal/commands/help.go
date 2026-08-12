// Copyright 2025 TechDivision GmbH
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

package commands

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/valet-sh/cli/internal/style"
)

// ErrorPrefix returns a styled "✘ msg" string for Go-layer error output.
// Used by root RunE (unknown command) and validation errors.
func ErrorPrefix(msg string) string {
	return style.Red(os.Stderr, "✘ "+msg)
}

func isAnsibleCommand(cmd *cobra.Command) bool {
	return cmd.Annotations["playbook"] != "" || cmd.Annotations["playbook-group"] != ""
}

func printCommandList(w io.Writer, header string, cmds []*cobra.Command) {
	_, _ = fmt.Fprintln(w, style.Blue(w, header))
	for _, sub := range cmds {
		if !sub.IsAvailableCommand() {
			continue
		}
		padding := strings.Repeat(" ", max(1, 20-len(sub.Name())))
		_, _ = fmt.Fprintf(w, "  %s%s%s\n",
			style.Green(w, sub.Name()),
			padding,
			sub.Short,
		)
	}
	_, _ = fmt.Fprintln(w)
}

// SetHelpFormatter installs a custom help renderer on the root command that
// cascades to all subcommands. Section headers get the blue ▶ prefix used
// by the Ansible callback's play_start output; command names are green.
func SetHelpFormatter(root *cobra.Command) {
	fn := helpFunc()
	root.SetHelpFunc(fn)
}

func helpFunc() func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, _ []string) {
		w := cmd.OutOrStdout()

		// Long description (or short if no long).
		desc := cmd.Long
		if desc == "" {
			desc = cmd.Short
		}
		if desc != "" {
			_, _ = fmt.Fprintln(w, desc)
			_, _ = fmt.Fprintln(w)
		}

		// Usage line.
		if cmd.Runnable() || cmd.HasAvailableSubCommands() {
			_, _ = fmt.Fprintln(w, style.Blue(w, "▶ Usage"))
			if cmd.Runnable() {
				_, _ = fmt.Fprintf(w, "  %s\n", cmd.UseLine())
			}
			if cmd.HasAvailableSubCommands() {
				_, _ = fmt.Fprintf(w, "  %s [options] [command] [arguments]\n", cmd.CommandPath())
			}
			_, _ = fmt.Fprintln(w)
		}

		if cmd.HasAvailableSubCommands() {
			var ansibleCmds, cliCmds []*cobra.Command
			for _, sub := range cmd.Commands() {
				if !sub.IsAvailableCommand() {
					continue
				}
				if isAnsibleCommand(sub) {
					ansibleCmds = append(ansibleCmds, sub)
				} else {
					cliCmds = append(cliCmds, sub)
				}
			}

			if len(ansibleCmds) > 0 && len(cliCmds) > 0 {
				printCommandList(w, "▶ Commands", ansibleCmds)
				printCommandList(w, "▶ CLI Commands", cliCmds)
			} else {
				printCommandList(w, "▶ Available Commands", cmd.Commands())
			}
		}

		// Flags.
		flags := cmd.LocalFlags()
		if cmd.HasAvailableLocalFlags() {
			_, _ = fmt.Fprintln(w, style.Blue(w, "▶ Flags"))
			_, _ = fmt.Fprintln(w, flags.FlagUsages())
		}

		// Inherited flags (only shown if there are any beyond help).
		if cmd.HasAvailableInheritedFlags() {
			_, _ = fmt.Fprintln(w, style.Blue(w, "▶ Global Flags"))
			_, _ = fmt.Fprintln(w, cmd.InheritedFlags().FlagUsages())
		}

		// Hint line.
		if cmd.HasAvailableSubCommands() {
			_, _ = fmt.Fprintf(w, "%s\n",
				style.Blue(w, fmt.Sprintf(`Use "%s [command] --help" for more information about a command.`, cmd.CommandPath())),
			)
		}
	}
}
