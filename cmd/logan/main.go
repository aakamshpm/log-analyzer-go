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

	file, err := os.Open(filepath) // the log file content will be still on disk even after open; the 'file' is a pointer to disk value

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	// we create a scanner buffer which holds chunk of bytes (lines) in memory from the file descriptor given
	scanner := bufio.NewScanner(file)
	totalLines := 0

	// Scan() returns true until a new line is read, it returns false if EOF/error is occured
	// it checks for /n new line everytime to load into the chunk as the current token
	// when /n is encountered, the data in that line becomes the current token
	// when encountered EOF, Scan() returns false
	for scanner.Scan() {
		totalLines++
	}

	// after scan, the Scan() returns false.
	// the false value may indicate either EOF or error
	// hence we need to check for errors using scanner.Err()
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total lines: %d\n", totalLines)
}
