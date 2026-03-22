package users

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type fakeStore struct {
	created []CreateInput
	users   []User
}

func (f *fakeStore) Create(_ context.Context, input CreateInput) (User, error) {
	f.created = append(f.created, input)
	return User{
		ID:        1,
		Name:      input.Name,
		Email:     input.Email,
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (f *fakeStore) List(_ context.Context, limit int) ([]User, error) {
	if limit > len(f.users) {
		limit = len(f.users)
	}
	return f.users[:limit], nil
}

func TestCreateUserReturnsCreated(t *testing.T) {
	store := &fakeStore{}
	h := NewHandlerWithStore(store)

	body := bytes.NewBufferString(`{"name":"Mylen","email":"mylen@example.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/users", body)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, rec.Code)
	}

	if len(store.created) != 1 {
		t.Fatalf("expected one create call, got %d", len(store.created))
	}
}

func TestListUsersReturnsOK(t *testing.T) {
	store := &fakeStore{
		users: []User{
			{ID: 2, Name: "Ana", Email: "ana@example.com", CreatedAt: time.Now().UTC()},
			{ID: 1, Name: "Bia", Email: "bia@example.com", CreatedAt: time.Now().UTC()},
		},
	}
	h := NewHandlerWithStore(store)

	req := httptest.NewRequest(http.MethodGet, "/users?limit=2", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}

	var got []User
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(got) != 2 {
		t.Fatalf("expected 2 users, got %d", len(got))
	}
}

func TestListUsersInvalidLimitReturnsBadRequest(t *testing.T) {
	store := &fakeStore{}
	h := NewHandlerWithStore(store)

	req := httptest.NewRequest(http.MethodGet, "/users?limit=0", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, rec.Code)
	}
}
