package user

import (
	"net/http"

	createUser "github.com/gabmenezesdev/golang-login-boilerplate/src/user/create-user"
	getOneUser "github.com/gabmenezesdev/golang-login-boilerplate/src/user/get-one-user"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createUser.CreateUserController(w, r)
	}

	if r.Method == "GET" {
		getOneUser.GetOneUserController(w, r)
	}

	if r.Method == "PUT" {
		getOneUser.GetOneUserController(w, r)
	}

	if r.Method == "DELETE" {
		getOneUser.GetOneUserController(w, r)
	}

	w.WriteHeader(http.StatusNotFound)
}

