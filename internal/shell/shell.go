package shell

import "os/exec"

// ExistCmd check whether command already installed in your system.
func ExistCmd(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
