package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	filename := args[1]

	fmt.Printf("Filename : %s\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to read file %s\n", filename)
		os.Exit(1)
	}
	f := make([]byte, 32*1024)
	file.Read(f)
	fmt.Println(string(f))
}
