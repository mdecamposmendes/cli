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

package setup

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/valet-sh/cli/internal/style"
)

func InstallLinuxDependencies() error {

	fmt.Printf("%s\n", style.Info(os.Stdout, "Running apt-get update"))

	cmd := exec.Command("sudo", "apt-get", "update")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to update apt: %w", err)
	}

	fmt.Printf("%s\n", style.Info(os.Stdout, "Installing dependencies"))

	cmd = exec.Command("sudo", "apt-get", "install", "-y", "git", "curl", "tar")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install dependencies: %w", err)
	}

	return nil
}
