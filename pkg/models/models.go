package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord error when user tries to find a record that
	// doesn't exist.
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials error when a user tries to login with
	// an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail error when a user tries to signup with an
	// email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet type
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User type
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
