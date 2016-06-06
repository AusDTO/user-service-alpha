package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	username       string
	hashedPassword []byte
}

func NewUser(username, password string) (*User, error) {
	// validate password length/complextity
	user := &User{
		username: username,
	}
	hashedPass, err := generatePassword(password)
	if err != nil {
		return user, err
	}
	user.hashedPassword = hashedPass
	return user, nil
}

func ExistingUser(username, clearPassword string, hashedPassword []byte) (*User, error) {
	user := &User{
		username: username,
	}
	if validPassword(hashedPassword, clearPassword) {
		user.hashedPassword = hashedPassword
		return user, nil
	}
	return user, fmt.Errorf("Incorrect username and password")
}

func generatePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func validPassword(hashedPassword []byte, clearPassword string) bool {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(clearPassword)) == nil
}
