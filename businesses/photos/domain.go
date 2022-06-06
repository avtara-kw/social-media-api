package photos

import (
	"github.com/avtara-kw/social-media-api/businesses/users"
	"time"
)

type Domain struct {
	ID        uint
	Title     string
	Caption   string
	PhotoURL  string
	UserID    int
	Users     users.Domain
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Service interface {
	Post(data *Domain) (*Domain, error)
	GetAll() ([]Domain, error)
	Update(data *Domain, photoID string) (*Domain, error)
	Delete(photoID, userID int) error
}

type Repository interface {
	Store(data *Domain) (*Domain, error)
	GetAll() ([]Domain, error)
	DeleteByID(photoID int) error
	Update(data *Domain, photoID string) (*Domain, error)
	GetByID(ID int) (*Domain, error)
}
