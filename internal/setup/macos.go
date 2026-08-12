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
