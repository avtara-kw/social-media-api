package comments

import (
	"github.com/avtara-kw/social-media-api/businesses/comments"
	"time"
)

type Comments struct {
	ID        uint
	Message   string
	PhotoID   int
	UserID    int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func toDomain(rec *Comments) *comments.Domain {
	return &comments.Domain{
		ID:        rec.ID,
		Message:   rec.Message,
		PhotoID:   rec.PhotoID,
		UserID:    rec.UserID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain *comments.Domain) *Comments {
	return &Comments{
		ID:        domain.ID,
		Message:   domain.Message,
		PhotoID:   domain.PhotoID,
		UserID:    domain.UserID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
