package getOneUser

import (
	"encoding/json"
	"net/http"
)

type User struct {
    Name    string `json:"name"`
    Surname string `json:"surname"`
    Email   string `json:"email"`
    Cpf     string `json:"cpf"`
}

func GetOneUserController(w http.ResponseWriter, r *http.Request) {
	var got User

	json.NewEncoder(w).Encode(got)
}
