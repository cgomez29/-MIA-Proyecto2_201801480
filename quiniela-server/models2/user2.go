package models2

type ADMIN struct {
	Idadmin uint `json:"idAdmin"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

