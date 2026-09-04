package main

import (
	"os"
	"os/exec"
	"syscall"
)

func createRoot(path string) {
	cmd := exec.Command(path)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// rootfs
	syscall.Mount("/root/bundle/rootfs",
		"/root/bundle/rootfs",
		"",
		syscall.MS_BIND|syscall.MS_REC,
		"")

	// pivot root
	syscall.PivotRoot("/root/bundle/rootfs", "./oldroot")
	// syscall.Chroot("/home/rashu/bundle/rootfs")
	os.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	panic(cmd.Run())
}

func root() {
	// reexec pattern
	// cmd := exec.Command("/proc/self/exe", append([]string{"child"}, "/bin/bash")...)

	// takes the process
	cmd := exec.Command("/proc/self/exec", append([]string{"createRoot"}, "/bin/bash")...)

	// gets connected to the given file

	// it requires stdin to take the arguments
	// from the bash shell

	//defines their standard input, output and error
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// defines attr for the cmd defined
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	// child()

	// ps cant see new process if its in the older process
	// but with a new mount space it can see the new ps
	panic(cmd.Run())
}

func main() {

	switch os.Args[1] {
	case "root":
		root()
	case "createRoot":
		createRoot(os.Args[2])
	}
}
