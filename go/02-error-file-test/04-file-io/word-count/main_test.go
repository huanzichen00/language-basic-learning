package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestCountTextStats(t *testing.T) {
	testCases := []struct {
		name      string
		content   string
		wantStats TextStats
	}{
		{
			name:    "two lines",
			content: "hello go\nthis is a test",
			wantStats: TextStats{
				Lines: 2,
				Words: 6,
				Chars: 22,
			},
		},
		{
			name:    "empty file",
			content: "",
			wantStats: TextStats{
				Lines: 0,
				Words: 0,
				Chars: 0,
			},
		},
		{
			name:    "multiple spaces",
			content: "hello   go\n",
			wantStats: TextStats{
				Lines: 1,
				Words: 2,
				Chars: 10,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "notes.txt")

			err := os.WriteFile(path, []byte(tc.content), 0644)
			if err != nil {
				t.Fatalf("write test file failed: %v", err)
			}

			gotStats, err := countTextStats(path)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if gotStats != tc.wantStats {
				t.Fatalf("expected %+v, got %+v", tc.wantStats, gotStats)
			}
		})
	}
}

func TestCountTextStatsFileNotExist(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "missing.txt")

	_, err := countTextStats(path)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected os.ErrNotExist, got %v", err)
	}
}
