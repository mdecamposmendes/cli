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

func InstallMacOSDependencies() error {

	fmt.Printf("%s\n", style.Green(os.Stdout, "Downloading Homebrew"))

	cmd := exec.Command("curl", "-fsSL", "https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh", "-o", "/tmp/homebrew_install.sh")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download homebrew install script: %w", err)
	}

	fmt.Printf("%s\n", style.Green(os.Stdout, "Installing Homebrew"))

	cmd = exec.Command("/bin/bash", "/tmp/homebrew_install.sh")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install homebrew: %w", err)
	}

	return nil
}
