package main

import (
	"fmt"
	"os"
)

const usage = "Usage: analyze <file-path> --json"

func parseArgs(args []string) (filePath string, jsonMode bool, err error) {
	if len(os.Args) < 3 || len(os.Args) > 4 || os.Args[1] != "analyze" {
		return "", false, fmt.Errorf(usage)
	}

	if len(os.Args) == 4 {
		if os.Args[3] != "--json" {
			return "", false, fmt.Errorf(usage)
		}
		jsonMode = true
	}
	return args[2], jsonMode, nil
}
