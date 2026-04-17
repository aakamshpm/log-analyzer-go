package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Summary struct {
	TotalLines int `json:"total_lines"`
	Warn       int `json:"warn"`
	Info       int `json:"info"`
	Error      int `json:"error"`
}

func main() {
	filePath, jsonMode, err := parseArgs(os.Args)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	file, err := os.Open(filePath) // the log file content will be still on disk even after open; the 'file' is a pointer/reader to disk value

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	summary, err := Analyzer(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
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
		fmt.Printf("Total lines: %d\n", summary.TotalLines)
		fmt.Printf("INFO: %d\n", summary.Info)
		fmt.Printf("WARN: %d\n", summary.Warn)
		fmt.Printf("ERROR: %d\n", summary.Error)
	}
}
