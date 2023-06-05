package main

import (
	"fmt"
	"log"
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
	fmt.Printf("Running %v as process %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	syscall.Chroot()
	if err := cmd.Run(); err != nil {
		log.Fatal("Error: ", err)
	}

}

func child() {

}
