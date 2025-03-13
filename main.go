package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"path/filepath"
   	"strconv"
 	"io/ioutil"
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

	// cgroups
	cg()

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

  func cg() {
   	cgroupRoot := "/sys/fs/cgroup"
   	containerCgroup := filepath.Join(cgroupRoot, "container")
   
   	// 1. Delete existing cgroup (if any)
   	os.RemoveAll(containerCgroup)
   
   	// 2. Create cgroup directory
   	if err := os.Mkdir(containerCgroup, 0755); err != nil {
   		panic(fmt.Sprintf("Failed to create cgroup: %v", err))
   	}
   
   		// Set memory limit (100MB)
   	must(ioutil.WriteFile(
   		filepath.Join(containerCgroup, "memory.max"),
   		[]byte("100M"), // 100 megabytes
   		0644,
   	))
   
   	// Set CPU limit 
   	must(ioutil.WriteFile(
   		filepath.Join(containerCgroup, "cpu.max"),
   		[]byte("50000 100000"), // 50ms quota, 100ms period
   		0644,
   	))
   
   	// 5. Add current process to cgroup
   	pid := os.Getpid()
   	if err := ioutil.WriteFile(
   		filepath.Join(containerCgroup, "cgroup.procs"),
   		[]byte(strconv.Itoa(pid)),
   		0644,
   	); err != nil {
   		panic(fmt.Sprintf("Failed to add process: %v", err))
   	}
   }

    // Helper function to check if a string contains a substring
   func contains(s, substr string) bool {
   	for i := 0; i < len(s)-len(substr)+1; i++ {
   		if s[i:i+len(substr)] == substr {
   			return true
   		}
   	}
   	return false
   }