package main

import (
	"os"
	"fmt"
	"strings"
	"database/sql"

	_ "github.com/lib/pq"
)

type UserDB struct {
	db *sql.DB
}

func GetUserDB() (*UserDB, error) {
	userDb := &UserDB{}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		return userDb, fmt.Errorf("DATABASE_URL empty")
	}
	// some random reconnect thing on the end which breaks everything. GEt RID
	dbUrl = strings.Replace(dbUrl, "reconnect=true", "", 1)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return userDb, err
	}
	userDb.db = db
	return userDb, nil
}

func (u *UserDB) getUser(username, clearPassword string) (*User, error) {
	var (
		user           string
		hashedPassword string
	)
	err := u.db.QueryRow("SELECT email, password FROM users WHERE email=$1", username).Scan(&user, &hashedPassword)
	if err != nil {
		return &User{}, err
	}

	return ExistingUser(user, clearPassword, []byte(hashedPassword))
}

func (u *UserDB) createUser(user *User) error {
	_, err := u.db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.username, string(user.hashedPassword))
	return err
}
