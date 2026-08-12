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

package updater

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/valet-sh/cli/internal/ansible"
)

const (
	serviceFile = "/usr/local/valet-sh/etc/services.yml"
	bundlesFile = "/usr/local/valet-sh/etc/bundles.yml"
)

var migrationText = `
┌────────────────────────────────────────────────────────────────────┐
│ valet.sh — Migration Required (v2.x Detected)                      │
└────────────────────────────────────────────────────────────────────┘

  WARNING: Existing services will be uninstalled and NO database
  dumps will be generated automatically.

  Please backup all critical database and application data.

  Post-Migration Environment:
  • Base Services Installed: Nginx, Dnsmasq, Mailpit, Container Runtime
    (Podman/Apple Container)
  • Optional Services Removed: PHP, MySQL, MariaDB, RabbitMQ, etc.
    (Must be re-installed on demand)

  For migration support and manual backup steps, visit:
  https://valet.sh/3.x/how-to-articles/migrating-from-2.x-to-3.x


`

func CheckMigration(repoDir string) error {
	if PlaybookBranch != "3.x" {
		return nil
	}

	_, err := os.Stat(serviceFile)
	serviceExists := err == nil
	_, err = os.Stat(bundlesFile)
	bundleExists := err == nil

	if serviceExists || (serviceExists && bundleExists) {
		fmt.Printf("%s", migrationText)

		fmt.Print("Type 'migrate' to proceed, or press [Enter] to cancel/skip: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		trimmedInput := strings.TrimSpace(input)

		if trimmedInput == "" || trimmedInput != "migrate" {
			fmt.Println("\n Migration canceled. Do you want to switch to the 2.x branch instead? (y/n): ")
			confirmInput, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("failed to read input: %w", err)
			}

			trimmedConfirmInput := strings.TrimSpace(confirmInput)

			if trimmedConfirmInput == "y" || trimmedConfirmInput == "Y" {
				if err := setChannel("2.x", repoDir); err != nil {
					return fmt.Errorf("failed to switch to 2.x branch: %w", err)
				}
			} else {
				fmt.Println("Migration canceled")
				os.Exit(0)
			}

			return nil
		}

		ansible.SetVar("valet_migrate", true)
	}

	return nil
}
