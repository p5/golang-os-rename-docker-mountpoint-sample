package main

import (
	"fmt"
	"os"
)

func main() {
	// Accept a file name as a command line argument
	fileName := os.Args[1]
	fmt.Println("File name:", fileName)

	// Create a new file
	newfile, err := os.Create("newfile.txt")
	if err != nil {
		os.Exit(1)
	}
	defer newfile.Close()
	fmt.Println("New file created:", newfile.Name())

	// Rename the new file to the original file name
	if err := os.Rename(newfile.Name(), fileName); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("File renamed to:", fileName)
}
