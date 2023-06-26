package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// go run coal main.go
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("invalid command")
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	
	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v as process %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot("rootfs"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "/proc", "proc", 0, ""))

	must(cmd.Run())

	must(syscall.Unmount("/proc", 0))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}