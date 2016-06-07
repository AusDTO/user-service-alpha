package main

import (
	"os"
	"strings"
	"database/sql"
	"io/ioutil"

	_ "github.com/lib/pq"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		panic("DATABASE_URL empty")
	}
	// some random reconnect thing on the end which breaks everything. GEt RID
	dbUrl = strings.Replace(dbUrl, "reconnect=true", "", 1)
	bytes, err := ioutil.ReadFile("./users.sql")
	if err != nil {
		panic(err)
	}

	fileString := string(bytes)
	fileString = strings.Replace(fileString, "\n", " ", -1)
	cmds := strings.Split(fileString, ";")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}

	for _, cmd := range cmds {
		_, err = db.Exec(cmd)
		if err != nil {
			panic(err)
		}
	}
}
