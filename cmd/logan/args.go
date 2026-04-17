package main

import (
	"fmt"
)

const usage = "Usage: analyze <file-path> --json"

func parseArgs(args []string) (filePath string, jsonMode bool, err error) {
	if len(args) < 3 || len(args) > 4 || args[1] != "analyze" {
		return "", false, fmt.Errorf(usage)
	}

	if len(args) == 4 {
		if args[3] != "--json" {
			return "", false, fmt.Errorf(usage)
		}
		jsonMode = true
	}
	return args[2], jsonMode, nil
}
