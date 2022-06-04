package users

import "github.com/avtara-kw/social-media-api/businesses/users"

type RequestUserRegistration struct {
	Age      string `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
	Username string `json:"username" binding:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (rec *RequestUserRegistration) UserRegistrationToDomain() *users.Domain {
	return &users.Domain{
		Age:      rec.Age,
		Email:    rec.Email,
		Password: rec.Password,
		Username: rec.Username,
	}
}
