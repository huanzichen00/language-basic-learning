package main

import (
	"errors"
	"testing"
)

func TestNewService(t *testing.T) {
	repo := NewMemoryRepository(map[int]User{
		1: {
			ID:   1,
			Name: "Tom",
		},
	})

	service, err := NewService(repo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if service == nil {
		t.Fatalf("expected service, got nil")
	}
}

func TestNewServiceNilRepo(t *testing.T) {
	service, err := NewService(nil)
	if err == nil {
		t.Fatalf("expected err, got nil")
	}

	if service != nil {
		t.Fatal("expected nil service")
	}
}

func TestServiceGetUserName(t *testing.T) {
	repo := NewMemoryRepository(map[int]User{
		1: {
			ID:   1,
			Name: "Tom",
		},
		2: {
			ID:   2,
			Name: "Rose",
		},
	})

	service, err := NewService(repo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	testCases := []struct {
		name      string
		id        int
		wantName  string
		wantError bool
	}{
		{
			name:      "found Tom",
			id:        1,
			wantName:  "Tom",
			wantError: false,
		},
		{
			name:      "found Rose",
			id:        2,
			wantName:  "Rose",
			wantError: false,
		},
		{
			name:      "not found",
			id:        99,
			wantName:  "",
			wantError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotName, err := service.GetUserName(tc.id)

			if tc.wantError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}

				if !errors.Is(err, ErrUserNotFound) {
					t.Fatalf("expected ErrUserNotFound, got %v", err)
				}

				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if gotName != tc.wantName {
				t.Fatalf("expected %s, got %s", tc.wantName, gotName)
			}
		})
	}
}
