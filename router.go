package main

import (
	"net/http"

	user "github.com/gabmenezesdev/golang-login-boilerplate/src/user"
)

func InitRoutes() {
	http.HandleFunc("/user", user.UsersHandler)
}

