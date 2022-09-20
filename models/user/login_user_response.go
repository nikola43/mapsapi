package models

type LoginUserResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
