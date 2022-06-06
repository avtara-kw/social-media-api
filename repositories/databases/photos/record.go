package photos

import (
	"github.com/avtara-kw/social-media-api/businesses/photos"
	"time"
)

type Photos struct {
	ID        uint
	Title     string
	Caption   string
	PhotoURL  string
	UserID    int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func toDomain(rec *Photos) *photos.Domain {
	return &photos.Domain{
		ID:        rec.ID,
		Title:     rec.Title,
		Caption:   rec.Caption,
		PhotoURL:  rec.PhotoURL,
		UserID:    rec.UserID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain *photos.Domain) *Photos {
	return &Photos{
		ID:        domain.ID,
		Title:     domain.Title,
		Caption:   domain.Caption,
		PhotoURL:  domain.PhotoURL,
		UserID:    domain.UserID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
