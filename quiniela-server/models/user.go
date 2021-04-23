package models

type ADMIN struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}