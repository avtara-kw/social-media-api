package users

import "time"

type Domain struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	Age       int
	Token     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Service interface {
	Registration(data *Domain) (*Domain, error)
	Login(data *Domain) (*Domain, error)
	Delete(ID string) error
}

type Repository interface {
	Store(data *Domain) (*Domain, error)
	GetByEmail(email string) (*Domain, error)
	DeleteByID(ID string) error
}
