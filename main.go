package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)
 
 func main() {
 	if len(os.Args) < 2 {
 		panic("Usage: <command>")
 	}
 
 	switch os.Args[1] {
 	case "run":
		run()
	case "child":
		child()
	default:
 		panic("Bad command")
 	}
 }

 func run() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
 	cmd.Stdout = os.Stdout
 	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
 		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
 	}


 	if err := cmd.Run(); err != nil {
 		fmt.Printf("Error running command: %v\n", err)
 	}
 }

 func child() {
 	fmt.Printf("Running %v as %d in child\n", os.Args[2:], os.Getpid())

	// Set hostname
 	if err := syscall.Sethostname([]byte("oracleContainer")); err != nil {
 		fmt.Printf("Failed to set hostname: %v\n", err)
 		os.Exit(1)
 	}

	// Change root filesystem
 	rootfs := "/tmp/oracletalk"
 	if err := syscall.Chroot(rootfs); err != nil {
 		fmt.Printf("Failed to chroot: %v\n", err)
 		os.Exit(1)
 	}
 	if err := os.Chdir("/"); err != nil {
 		fmt.Printf("Failed to change directory: %v\n", err)
 		os.Exit(1)
 	}

	must(syscall.Mount("proc", "proc", "proc", 0, ""))

 
 	cmd := exec.Command(os.Args[2], os.Args[3:]...)
 	cmd.Stdin = os.Stdin
 	cmd.Stdout = os.Stdout
 	cmd.Stderr = os.Stderr
 
 	if err := cmd.Run(); err != nil {
 		fmt.Printf("Error running command in child: %v\n", err)
 	}
 }

 func must(err error) {
	if err != nil {
		panic(err)
	}
}