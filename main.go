package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Printf("Hello world!\nArguments: %v\n", args)
	} else {
		fmt.Println("Need args. Example: ./main.go arg1")
		os.Exit(1)
	}

}
