package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required, email"`
	Password  string `json:"password" validate:"required,mun=6,max=130"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUSerByID(id int) (*User, error)
	GetUSerByUSername(username string) (*User, error)
	CReateUSer(User) error
}
