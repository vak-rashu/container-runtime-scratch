package main

import (
	"os"
	"os/exec"
	"syscall"
)

func child() {
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	syscall.Mount("/home/rashu/bundle/rootfs",
		"/home/rashu/bundle/rootfs",
		"",
		syscall.MS_BIND|syscall.MS_REC,
		"")

	syscall.PivotRoot("/home/rashu/bundle/rootfs", "oldroot")
	os.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	cmd.Run()
}

func main() {
	cmd := exec.Command("/proc/self/exe/", append([]string{"child"}, "/bin/bash")...)

	// gets connected to the given file

	// it requires stdin to take the arguments
	// from the bash shell
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	child()

	// ps cant see new process if its in the older process
	// but with a new mount space it can see the new ps
	cmd.Run()
}
