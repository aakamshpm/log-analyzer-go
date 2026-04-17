package main

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestAnalyze(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Summary
	}{
		{
			name: "mixed levels",
			input: "" +
				"2026-04-16T10:00:01Z INFO service started\n" +
				"2026-04-16T10:00:05Z INFO listening on :8080\n" +
				"2026-04-16T10:01:10Z WARN cache miss\n" +
				"2026-04-16T10:03:20Z ERROR database timeout\n",
			want: Summary{
				TotalLines: 4,
				Info:       2,
				Warn:       1,
				Error:      1,
			},
		},
		{
			name:  "empty input",
			input: "",
			want: Summary{
				TotalLines: 0,
				Info:       0,
				Warn:       0,
				Error:      0,
			},
		},
		{
			name: "unknown level still counts total",
			input: "" +
				"2026-04-16T10:00:01Z DEBUG boot\n" +
				"random line without level\n",
			want: Summary{
				TotalLines: 2,
				Info:       0,
				Warn:       0,
				Error:      0,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := Analyzer(strings.NewReader(tc.input))
			if err != nil {
				t.Fatalf("Analyze() returned error: %v", err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Analyze() mismatch\n got: %+v\nwant: %+v", got, tc.want)
			}
		})
	}
}

type brokenReader struct{}

func (brokenReader) Read(_ []byte) (int, error) {
	return 0, errors.New("forced read error")
}
func TestAnalyze_ReadError(t *testing.T) {
	_, err := Analyzer(brokenReader{})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
