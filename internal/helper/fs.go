package helper

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func EnsureOwnedDir(path, root string) error {
	if err := mkdirAll(path); err != nil {
		return fmt.Errorf("failed to create %s: %w", path, err)
	}

	uid, gid, username, err := currentUser()
	if err != nil {
		return err
	}

	root = filepath.Clean(root)
	for dir := filepath.Clean(path); ; dir = filepath.Dir(dir) {
		if err := chown(dir, uid, gid, username); err != nil {
			return fmt.Errorf("failed to chown %s: %w", dir, err)
		}
		if dir == root || dir == filepath.Dir(dir) {
			break
		}
	}
	return nil
}

func mkdirAll(path string) error {
	if err := os.MkdirAll(path, 0o755); err != nil {
		if !os.IsPermission(err) {
			return err
		}
		return exec.Command("sudo", "mkdir", "-p", path).Run()
	}
	return nil
}

func chown(path string, uid, gid int, username string) error {
	if err := os.Chown(path, uid, gid); err != nil {
		if !os.IsPermission(err) {
			return err
		}
		return exec.Command("sudo", "chown", fmt.Sprintf("%s:%d", username, gid), path).Run()
	}
	return nil
}

func currentUser() (uid, gid int, username string, err error) {
	u, err := user.Current()
	if err != nil {
		return 0, 0, "", fmt.Errorf("could not determine current user: %w", err)
	}
	uid, err = strconv.Atoi(u.Uid)
	if err != nil {
		return 0, 0, "", fmt.Errorf("could not parse uid %q: %w", u.Uid, err)
	}
	gid, err = strconv.Atoi(u.Gid)
	if err != nil {
		return 0, 0, "", fmt.Errorf("could not parse gid %q: %w", u.Gid, err)
	}
	return uid, gid, u.Username, nil
}
