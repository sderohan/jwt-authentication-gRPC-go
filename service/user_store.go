package service

import (
	"errors"
	"sync"
)

// UserStore is an interface to store users
type UserStore interface {
	// Save saves a user to store
	Save(*User) error
	// Find finds a user by username
	Find(username string) (*User, error)
}

// InMemoryStore stores users in memory
type InMemoryUserStore struct {
	mutex sync.RWMutex
	// multiple users with same username cannot be stored
	users map[string]*User
}

// NewInMemoryUserStore returns a new in-memory user store
func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*User),
	}
}

// Save saves a user to the store
func (store *InMemoryUserStore) Save(user *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return errors.New("User already exists")
	}

	store.users[user.Username] = user.Clone()
	return nil
}

// Find finds a user by username
func (store *InMemoryUserStore) Find(username string) (*User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.users[username]
	if user == nil {
		return nil, nil
	}

	return user.Clone(), nil
}
