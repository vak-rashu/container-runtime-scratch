package main

// import (
// 	"fmt"
// 	"syscall"

// 	libseccomp "github.com/seccomp/libseccomp-golang"
// )

// var syscalls = []string{
// 	"rt_sigaction", "mkdirat", "clone", "mmap", "readlinkat", "futex", "rt_sigprocmask",
// 	"mprotect", "write", "sigaltstack", "gettid", "read", "open", "close", "fstat", "munmap",
// 	"brk", "access", "execve", "getrlimit", "arch_prctl", "sched_getaffinity", "set_tid_address", "set_robust_list",
// }

// func whitelist(syscalls []string) {
// 	filter, err := libseccomp.NewFilter(libseccomp.ActErrno.SetReturnCode(int16(syscall.EPERM)))
// 	if err != nil {
// 		fmt.Printf("Error running %v", err)
// 	}

// 	for _, element := range syscalls {
// 		fmt.Println(element)
// 		syscallID, err := libseccomp.GetSyscallFromName(element)
// 		if err != nil {
// 			panic(err)
// 		}
// 		filter.AddRule(syscallID, libseccomp.ActAllow)
// 	}

// 	filter.Load()
// }

// func main() {

// 	whitelist(syscalls)

// 	err := syscall.Mkdir("/home/rashu/devfest", 0755)
// 	if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("Folder created!!")
// 	}
// }

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"syscall"
// )

// // go run main.go run <cmd> <args>
// func main() {
// 	// run()
// 	switch os.Args[1] {
// 	case "run":
// 		run()
// 	case "child":
// 		child()
// 	default:
// 		panic("help")
// 	}
// }

// func run() {
// 	fmt.Printf("Running %v \n", os.Args[2:])

// 	// reexec pattern
// 	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

// 	// cmd := exec.Command("/bin/bash")
// 	fmt.Println(os.Getpid())
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	cmd.SysProcAttr = &syscall.SysProcAttr{
// 		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
// 		Unshareflags: syscall.CLONE_NEWNS,
// 	}

// 	must(cmd.Run())
// }

// func child() {
// 	fmt.Printf("Running %v \n", os.Args[2:])

// 	cmd := exec.Command(os.Args[2], os.Args[3:]...)
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	// must(syscall.Sethostname([]byte("container")))
// 	// must(syscall.Chroot("/home/liz/ubuntufs"))
// 	// must(os.Chdir("/"))
// 	// must(syscall.Mount("proc", "proc", "proc", 0, ""))
// 	// must(syscall.Mount("thing", "mytemp", "tmpfs", 0, ""))

// 	// cg()

// 	must(cmd.Run())

// 	// must(syscall.Unmount("proc", 0))
// 	// must(syscall.Unmount("thing", 0))
// }

// // func cg() {
// // 	cgroups := "/sys/fs/cgroup/"
// // 	pids := filepath.Join(cgroups, "pids")
// // 	os.MkdirAll(filepath.Join(pids, "liz"), 0755)
// // 	must(ioutil.WriteFile(filepath.Join(pids, "liz/pids.max"), []byte("20"), 0700))
// // 	// Removes the new cgroup in place after the container exits
// // 	must(ioutil.WriteFile(filepath.Join(pids, "liz/notify_on_release"), []byte("1"), 0700))
// // 	must(ioutil.WriteFile(filepath.Join(pids, "liz/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
// // }

// func must(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func main() {
// 	// cmd := exec.Command("ls", "-al")

// 	// cmd.Stdin = os.Stdin
// 	// cmd.Stdout = os.Stdout
// 	// cmd.Stderr = os.Stderr

// 	// argSlice := cmd.Args
// 	argSlice := os.Args
// 	fmt.Println(argSlice)
// }
