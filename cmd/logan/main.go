package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "analyze" {
		fmt.Fprintln(os.Stderr, "Usage: analyze <file-path>")
		os.Exit(2)
	}

	filepath := os.Args[2]

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalLines := 0

	for scanner.Scan() {
		totalLines++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("Total lines: %d\n", totalLines)
}
