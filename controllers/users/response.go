package users

import (
	"github.com/avtara-kw/social-media-api/businesses/users"
	"time"
)

type ResponseUserRegistration struct {
	ID       uint   `json:"id"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ResponseUserLogin struct {
	Token string `json:"token"`
}

type ResponseUserUpdate struct {
	ID        uint       `json:"id"`
	Age       int        `json:"age"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	UpdatedAt *time.Time `json:"date"`
}

func ResponseUserRegistrationFromDomain(domain *users.Domain) *ResponseUserRegistration {
	return &ResponseUserRegistration{
		Age:      domain.Age,
		Email:    domain.Email,
		ID:       domain.ID,
		Username: domain.Username,
	}
}

func ResponseUserLoginFromDomain(domain *users.Domain) *ResponseUserLogin {
	return &ResponseUserLogin{
		Token: domain.Token,
	}
}

func ResponseUserUpdateFromDomain(domain *users.Domain) *ResponseUserUpdate {
	return &ResponseUserUpdate{
		Age:       domain.Age,
		Email:     domain.Email,
		ID:        domain.ID,
		Username:  domain.Username,
		UpdatedAt: domain.UpdatedAt,
	}
}
