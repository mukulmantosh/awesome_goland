package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	// defer file.Close()

	// Read the entire content
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("File content:\n%s\n", string(content))
}
