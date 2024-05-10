package createUser

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
    Name    string `json:"name"`
    Surname string `json:"surname"`
    Email   string `json:"email"`
    Cpf     string `json:"cpf"`
}

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
        http.Error(w, "Failed to parse request body", http.StatusBadRequest)
        return
    }
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(reqBody)
}
