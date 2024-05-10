package user

type User  struct{
    Name    string `json:"name"`
    Surname string `json:"surname"`
    Email   string `json:"email"`
    Cpf     string `json:"cpf"`
	PhoneNumber     string `json:"phoneNumber"`
}