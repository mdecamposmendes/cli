package setup

import (
	"fmt"
	goruntime "runtime"

	"os"

	"github.com/valet-sh/cli/constants"
	"github.com/valet-sh/cli/internal/helper"
	"github.com/valet-sh/cli/internal/style"
	"github.com/valet-sh/cli/internal/updater"
)

func Setup() error {
	proceed, err := checkExistingInstallation()
	if err != nil {
		return err
	}
	if !proceed {
		return nil
	}

	if err := installDependencies(); err != nil {
		return err
	}

	return commonSetup(constants.VshBasePath)
}

func checkExistingInstallation() (bool, error) {
	if helper.PathExists(constants.VshBasePath) || helper.PathExists(constants.VshVenvPath) {
		fmt.Printf("%s\n", style.Yellow(os.Stdout, "valet-sh is already installed. Do you want to reinstall? (y/n): "))
		var response string
		fmt.Scanln(&response)
		if response != "y" {
			fmt.Printf("%s\n", style.Yellow(os.Stdout, "Setup cancelled"))
			return false, nil
		} else {
			fmt.Printf("%s\n", style.Yellow(os.Stdout, "Removing existing installation..."))
			if err := removeInstallation(); err != nil {
				return false, fmt.Errorf("failed to remove existing installation: %w", err)
			}
			fmt.Printf("%s\n", style.Green(os.Stdout, "Existing installation removed"))
		}
	}

	return true, nil
}

func commonSetup(repoDir string) error {
	fmt.Printf("%s\n", style.Green(os.Stdout, "Installing Valet-sh..."))

	if _, err := updater.EnsurePlaybooks(repoDir, updater.PlaybookRepo, updater.PlaybookBranch); err != nil {
		return fmt.Errorf("failed to install playbooks: %w", err)
	}
	if _, err := updater.EnsureRuntime(repoDir); err != nil {
		return fmt.Errorf("failed to install runtime: %w", err)
	}

	fmt.Printf("%s\n", style.Green(os.Stdout, "Valet-sh setup complete"))
	return nil
}

func installDependencies() error {
	switch goruntime.GOOS {
	case "linux":
		return InstallLinuxDependencies()
	case "darwin":
		return InstallMacOSDependencies()
	default:
		return fmt.Errorf("unsupported operating system: %s", goruntime.GOOS)
	}
}

func removeInstallation() error {
	if _, err := os.Stat(constants.VshAnsibleFactsFile); err == nil {
		if err := os.Remove(constants.VshAnsibleFactsFile); err != nil {
			return fmt.Errorf("failed to remove ansible facts file: %w", err)
		}
	}

	if _, err := os.Stat(constants.VshBasePath); err == nil {
		if err := os.RemoveAll(constants.VshBasePath); err != nil {
			return fmt.Errorf("failed to remove existing repository: %w", err)
		}
	}

	if _, err := os.Stat(constants.VshVenvPath); err == nil {
		if err := os.RemoveAll(constants.VshVenvPath); err != nil {
			return fmt.Errorf("failed to remove existing virtual environment: %w", err)
		}
	}

	return nil
}
