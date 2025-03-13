package main

import (
	"fmt"
	"os"
	"os/exec"
)
 
 func main() {
 	if len(os.Args) < 2 {
 		panic("Usage: <command>")
 	}
 
 	switch os.Args[1] {
 	case "run":
		run()
	default:
 		panic("Bad command")
 	}
 }

 func run() {
 	fmt.Printf("Running %v\n", os.Args[2:])
 
 	cmd := exec.Command(os.Args[2], os.Args[3:]...)
 	cmd.Stdin = os.Stdin
 	cmd.Stdout = os.Stdout
 	cmd.Stderr = os.Stderr
 
 	if err := cmd.Run(); err != nil {
 		fmt.Printf("Error running command: %v\n", err)
 	}
 }