package users

import (
	"github.com/avtara-kw/social-media-api/businesses/users"
	"time"
)

type Users struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	Age       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func toDomain(rec *Users) *users.Domain {
	return &users.Domain{
		ID:        rec.ID,
		Username:  rec.Username,
		Email:     rec.Email,
		Password:  rec.Password,
		Age:       rec.Age,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain *users.Domain) *Users {
	return &Users{
		ID:        domain.ID,
		Username:  domain.Username,
		Email:     domain.Email,
		Password:  domain.Password,
		Age:       domain.Age,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
