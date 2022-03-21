package shell

import (
	"os"
	"os/exec"
)

// ExistCmd check whether command already installed in your system.
func ExistCmd(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// CurrentShell return $SHELL value
func CurrentShell() string {
	return os.Getenv("SHELL")
}
