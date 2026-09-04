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
	// syscall.Mount("/home/rashu/bundle/rootfs",
	// 	"/home/rashu/bundle/rootfs",
	// 	"",
	// 	syscall.MS_BIND|syscall.MS_REC,
	// 	"")

	// pivot root
	// syscall.PivotRoot("/home/rashu/bundle/rootfs", "./oldroot")
	syscall.Chroot("/home/rashu/bundle/rootfs")
	// syscall.Chroot("/home/rashu/bundle/rootfs")
	os.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	panic(cmd.Run())
}

func run() {
	// reexec pattern
	// cmd := exec.Command("/proc/self/exe", append([]string{"child"}, "/bin/bash")...)

	// takes the process
	cmd := exec.Command("/proc/self/exe", append([]string{"createRoot"}, "/bin/sh")...)

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
	case "run":
		run()
	case "createRoot":
		// fmt.Println("whatt")
		createRoot(os.Args[2])
	}
}
