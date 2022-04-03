package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username       string
	HashedPassword string
	Role           string
}

// NewUser returns a new user
func NewUser(username string, password string, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password %w", err)
	}

	user := &User{
		Username:       username,
		HashedPassword: string(hashedPassword),
		Role:           role,
	}
	return user, nil
}

// Clone return a user copy of this user
func (user *User) Clone() *User {
	return &User{
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
	}
}
