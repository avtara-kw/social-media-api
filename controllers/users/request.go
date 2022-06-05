package users

import "github.com/avtara-kw/social-media-api/businesses/users"

type RequestUserRegistration struct {
	Age      int    `json:"age" binding:"required,gt=8"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
	Username string `json:"username" binding:"required"`
}

type RequestUserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (rec *RequestUserRegistration) ToDomain() *users.Domain {
	return &users.Domain{
		Age:      rec.Age,
		Email:    rec.Email,
		Password: rec.Password,
		Username: rec.Username,
	}
}

func (rec *RequestUserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    rec.Email,
		Password: rec.Password,
	}
}
