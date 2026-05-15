package user

import "testing"

func TestServiceGetUserName(t *testing.T) {
	repo := NewMemoryUserRepository(map[int]User{
		1: {
			ID:   1,
			Name: "Tom",
		},
		2: {
			ID:   2,
			Name: "Rose",
		},
	})

	service := NewUserService(repo)

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
