package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func resetUsers() {
	users = []User{
		{
			ID:   1,
			Name: "Tom",
		},
		{
			ID:   2,
			Name: "Rose",
		},
	}
}

func newRequest(method string, path string, body string) *http.Request {
	req := httptest.NewRequest(method, path, strings.NewReader(body))

	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}

	return req
}

func executeRequest(handler http.HandlerFunc, req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	handler(recorder, req)
	return recorder
}

func assertStatus(t *testing.T, recorder *httptest.ResponseRecorder, want int) {
	t.Helper()

	if recorder.Code != want {
		t.Fatalf("expected status %d, got %d", want, recorder.Code)
	}
}

func decodeBody[T any](t *testing.T, recorder *httptest.ResponseRecorder) T {
	t.Helper()

	var value T

	err := json.NewDecoder(recorder.Body).Decode(&value)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	return value
}

func TestUsersHandlerGet(t *testing.T) {
	resetUsers()

	req := newRequest(http.MethodGet, "/users", "")
	recorder := executeRequest(usersHandler, req)

	assertStatus(t, recorder, http.StatusOK)

	got := decodeBody[[]User](t, recorder)

	if len(got) != 2 {
		t.Fatalf("expected 2 users, got %d", len(got))
	}
}

func TestUserByIDHandlerGetFound(t *testing.T) {
	resetUsers()

	req := newRequest(http.MethodGet, "/users/1", "")
	recorder := executeRequest(userByIDHandler, req)

	assertStatus(t, recorder, http.StatusOK)

	got := decodeBody[User](t, recorder)

	if got.ID != 1 {
		t.Fatalf("expected user ID 1, got %d", got.ID)
	}

	if got.Name != "Tom" {
		t.Fatalf("expected Tom, got %s", got.Name)
	}
}

func TestUserByIDHandlerGetNotFound(t *testing.T) {
	resetUsers()

	req := newRequest(http.MethodGet, "/users/99", "")
	recorder := executeRequest(userByIDHandler, req)

	assertStatus(t, recorder, http.StatusNotFound)
}

func TestUsersHandlerPost(t *testing.T) {
	resetUsers()

	req := newRequest(http.MethodPost, "/users", `{"name":"Jack"}`)
	recorder := executeRequest(usersHandler, req)

	assertStatus(t, recorder, http.StatusCreated)

	got := decodeBody[User](t, recorder)

	if got.ID != 3 {
		t.Fatalf("expected user ID 3, got %d", got.ID)
	}

	if got.Name != "Jack" {
		t.Fatalf("expected Jack, got %s", got.Name)
	}
}

func TestUsersHandlerPostInvalidBody(t *testing.T) {
	resetUsers()

	req := newRequest(http.MethodPost, "/users", `{name:"Jack"}`)
	recorder := executeRequest(usersHandler, req)

	assertStatus(t, recorder, http.StatusBadRequest)
}

func TestUserByIDHandlerPut(t *testing.T) {
	resetUsers()

	req := newRequest(http.MethodPut, "/users/2", `{"name":"RoseUpdated"}`)
	recorder := executeRequest(userByIDHandler, req)

	assertStatus(t, recorder, http.StatusOK)

	got := decodeBody[User](t, recorder)

	if got.ID != 2 {
		t.Fatalf("expected user ID 2, got %d", got.ID)
	}
	if got.Name != "RoseUpdated" {
		t.Fatalf("expected RoseUpdated, got %s", got.Name)
	}
}

func TestUserByIDHandlerDelete(t *testing.T) {
	resetUsers()

	req := newRequest(http.MethodDelete, "/users/2", "")
	recorder := executeRequest(userByIDHandler, req)

	assertStatus(t, recorder, http.StatusNoContent)

	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}

	if users[0].ID != 1 {
		t.Fatalf("expected remaining user ID 1, got %d", users[0].ID)
	}
}
