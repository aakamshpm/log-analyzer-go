package main

import (
	"bufio"
	"io"
	"strings"
)

func Analyzer(r io.Reader) (Summary, error) {
	// we create a scanner buffer which holds chunk of bytes (lines) in memory from the file descriptor given
	scanner := bufio.NewScanner(r)
	summary := Summary{}

	// Scan() returns true until a new line is read, it returns false if EOF/error is occured
	// it checks for /n new line everytime to load into the chunk as the current token
	// when /n is encountered, the data in that line becomes the current token
	// when encountered EOF, Scan() returns false
	for scanner.Scan() {
		line := scanner.Text() // returns the current token/line read

		if strings.Contains(line, " INFO ") {
			summary.Info++
		} else if strings.Contains(line, " WARN ") {
			summary.Warn++
		} else if strings.Contains(line, " ERROR ") {
			summary.Error++
		}

		summary.TotalLines++
	}

	// after scan, the Scan() returns false.
	// the false value may indicate either EOF or error
	// hence we need to check for errors using scanner.Err()
	if err := scanner.Err(); err != nil {
		return Summary{}, err
	}

	return summary, nil

}
