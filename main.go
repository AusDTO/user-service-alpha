/**
This service acts as the gatekeeper to the user database
**/
package main

import (
	"net/http"
	"os"
)

func main() {
	db, err := GetUserDB()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		pass := r.FormValue("password")

		if username == "" || pass == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		switch r.URL.Path {
		case "/login":
			if _, err = db.getUser(username, pass); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusOK)
			return
		case "/create":
			user, err := NewUser(username, pass)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err = db.createUser(user); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.ListenAndServe(":"+port, nil)
}
