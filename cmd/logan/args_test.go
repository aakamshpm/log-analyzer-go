package main

import "testing"

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantPath  string
		wantJSON  bool
		expectErr bool
	}{
		{
			name:      "valid without json",
			args:      []string{"logan", "analyze", "testdata/sample.log"},
			wantPath:  "testdata/sample.log",
			wantJSON:  false,
			expectErr: false,
		},
		{
			name:      "valid with json",
			args:      []string{"logan", "analyze", "testdata/sample.log", "--json"},
			wantPath:  "testdata/sample.log",
			wantJSON:  true,
			expectErr: false,
		},
		{
			name:      "invalid command",
			args:      []string{"logan", "scan", "testdata/sample.log"},
			expectErr: true,
		},
		{
			name:      "missing file path",
			args:      []string{"logan", "analyze"},
			expectErr: true,
		},
		{
			name:      "invalid flag",
			args:      []string{"logan", "analyze", "testdata/sample.log", "--badflag"},
			expectErr: true,
		},
		{
			name:      "too many args",
			args:      []string{"logan", "analyze", "testdata/sample.log", "--json", "extra"},
			expectErr: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotPath, gotJSON, err := parseArgs(tc.args)
			if tc.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotPath != tc.wantPath {
				t.Fatalf("path mismatch: got %q want %q", gotPath, tc.wantPath)
			}
			if gotJSON != tc.wantJSON {
				t.Fatalf("jsonMode mismatch: got %v want %v", gotJSON, tc.wantJSON)
			}
		})
	}
}
