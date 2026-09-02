package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("/bin/bash")

	// gets connected to the given file

	// it requires stdin to take the arguments
	// from the bash shell
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	// ps cant see new process if its in the older process
	// but with a new mount space it can see the new ps
	cmd.Run()
}
