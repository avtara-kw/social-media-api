package users

import "github.com/avtara-kw/social-media-api/businesses/users"

type ResponseUserRegistration struct {
	ID       uint   `json:"id"`
	Age      string `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func FromDomain(domain *users.Domain) *ResponseUserRegistration {
	return &ResponseUserRegistration{
		Age:      domain.Age,
		Email:    domain.Email,
		ID:       domain.ID,
		Username: domain.Username,
	}
}
