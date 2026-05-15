package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestCountLines(t *testing.T) {
	testCases := []struct {
		name      string
		content   string
		wantCount int
	}{
		{
			name:      "three lines",
			content:   "first line\nsecond line\nthird line",
			wantCount: 3,
		},
		{
			name:      "one line",
			content:   "only one line",
			wantCount: 1,
		},
		{
			name:      "empty file",
			content:   "",
			wantCount: 0,
		},
		{
			name:      "trailing newline",
			content:   "first line\nsecond line\n",
			wantCount: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()

			path := filepath.Join(dir, "notes.txt")

			err := os.WriteFile(path, []byte(tc.content), 0644)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			count, err := countLines(path)
			if count != tc.wantCount {
				t.Fatalf("expected %d lines, got %d", tc.wantCount, count)
			}
		})
	}
}

func TestCountLinesFileNotExist(t *testing.T) {
	dir := t.TempDir()

	path := filepath.Join(dir, "missing.txt")

	_, err := countLines(path)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected os.ErrNotExist, got %v", err)
	}
}
