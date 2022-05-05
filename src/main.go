package main

import (
	"fmt"
	"os"
)

func pwd() {
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(cwd)
}

func shell() {
	var cmd string
	fmt.Printf("secsh$ ")
	fmt.Scanln(&cmd)
	if (cmd == "cd") {
		fmt.Printf("Under development!\n")		
	} else if (cmd == "exit") {
		os.Exit(0)	
	} else if (cmd == "pwd") {
		pwd()
	}
}

func main() {
	for {
		shell()
	}
}
