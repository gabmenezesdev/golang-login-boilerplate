package userController

import (
	"net/http"

	createUser "github.com/gabmenezesdev/golang-login-boilerplate/src/user/create-user-usecase"
	getOneUser "github.com/gabmenezesdev/golang-login-boilerplate/src/user/get-one-user-usecase"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createUser.CreateUserController(w, r)
	}

	if r.Method == "GET" {
		getOneUser.GetOneUserController(w, r)
	}

	w.WriteHeader(http.StatusNotFound)
}

