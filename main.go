package main
 
 import (
 	"fmt"
 	"os"
 )
 
 func main() {
 	if len(os.Args) < 2 {
 		panic("Usage: <command>")
 	}
 
 	switch os.Args[1] {
 	case "run":
 		fmt.Println("Running a new container...")
 	default:
 		panic("Bad command")
 	}
 }