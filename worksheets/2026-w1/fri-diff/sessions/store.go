// Package sessions is an in-memory session store.
package sessions

import (
	"errors"
	"sync"
	"time"
)

// ErrNotFound is returned when no session exists for a token.
var ErrNotFound = errors.New("sessions: not found")

// Session is one logged-in user's state.
type Session struct {
	Token   string
	UserID  string
	Created time.Time
	Scopes  []string
}

// Store keeps sessions in memory, safe for concurrent use.
type Store struct {
	mu   sync.RWMutex
	byID map[string]*Session
}

func New() *Store {
	return &Store{byID: make(map[string]*Session)}
}

// Put stores a session under its token.
func (s *Store) Put(sess Session) {
	s.mu.Lock()
	defer s.mu.Unlock()
	cp := sess
	cp.Scopes = append([]string(nil), sess.Scopes...)
	s.byID[sess.Token] = &cp
}

// Get returns a copy of the session for token.
func (s *Store) Get(token string) (Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sess, ok := s.byID[token]
	if !ok {
		return Session{}, ErrNotFound
	}
	cp := *sess
	cp.Scopes = append([]string(nil), sess.Scopes...)
	return cp, nil
}

// Delete removes a session. Deleting a missing token is not an error.
func (s *Store) Delete(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.byID, token)
}

// Len reports how many sessions are held.
func (s *Store) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.byID)
}
