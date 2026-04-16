package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Summary struct {
	TotalLines int `json:"total_lines"`
	Warn       int `json:"warn"`
	Info       int `json:"info"`
	Error      int `json:"error"`
}

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 || os.Args[1] != "analyze" {
		fmt.Fprintln(os.Stderr, "Usage: analyze <file-path> --json")
		os.Exit(2)
	}

	jsonMode := len(os.Args) == 4 && os.Args[3] == "--json"

	filepath := os.Args[2]

	file, err := os.Open(filepath) // the log file content will be still on disk even after open; the 'file' is a pointer to disk value

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	// we create a scanner buffer which holds chunk of bytes (lines) in memory from the file descriptor given
	scanner := bufio.NewScanner(file)
	infoCount := 0
	warnCount := 0
	errorCount := 0
	totalLines := 0

	// Scan() returns true until a new line is read, it returns false if EOF/error is occured
	// it checks for /n new line everytime to load into the chunk as the current token
	// when /n is encountered, the data in that line becomes the current token
	// when encountered EOF, Scan() returns false
	for scanner.Scan() {
		line := scanner.Text() // returns the current token/line read

		if strings.Contains(line, " INFO ") {
			infoCount++
		} else if strings.Contains(line, " WARN ") {
			warnCount++
		} else if strings.Contains(line, " ERROR ") {
			errorCount++
		}

		totalLines++
	}

	// after scan, the Scan() returns false.
	// the false value may indicate either EOF or error
	// hence we need to check for errors using scanner.Err()
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	summary := Summary{
		TotalLines: totalLines, Warn: warnCount, Info: infoCount, Error: errorCount,
	}

	if jsonMode {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ") //pretty json

		// converts the struct data into json format and prints it based on the indentation
		if err := enc.Encode(summary); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing JSON: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Total lines: %d\n", totalLines)
		fmt.Printf("INFO: %d\n", infoCount)
		fmt.Printf("WARN: %d\n", warnCount)
		fmt.Printf("ERROR: %d\n", errorCount)
	}
}
