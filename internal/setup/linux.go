package setup

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/valet-sh/cli/internal/style"
)

func InstallLinuxDependencies() error {

	fmt.Printf("%s\n", style.Green(os.Stdout, "Running apt-get update"))

	cmd := exec.Command("sudo", "apt-get", "update")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to update apt: %w", err)
	}

	fmt.Printf("%s\n", style.Green(os.Stdout, "Installing dependencies"))

	cmd = exec.Command("sudo", "apt-get", "install", "-y", "git", "curl", "tar")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install dependencies: %w", err)
	}

	return nil
}
