package users

import "time"

type Domain struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	Age       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Service interface {
	Registration(data *Domain) (*Domain, error)
}

type Repository interface {
	Store(data *Domain) (*Domain, error)
}
