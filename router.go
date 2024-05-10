package main

import (
	"fmt"
	"net/http"

	user "github.com/gabmenezesdev/golang-login-boilerplate/src/user"
)

func Router() {
	fmt.Println("aqui")
	http.HandleFunc("/users", user.UsersHandler)
}

