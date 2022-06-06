package social_medias

import (
	"github.com/avtara-kw/social-media-api/businesses/social_medias"
	"github.com/avtara-kw/social-media-api/businesses/users"
	"time"
)

type SocialMedias struct {
	ID             uint
	Name           string
	SocialMediaURL string
	UserID         int
	Users          users.Domain
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

func toDomain(rec *SocialMedias) *social_medias.Domain {
	return &social_medias.Domain{
		ID:             rec.ID,
		Name:           rec.Name,
		SocialMediaURL: rec.SocialMediaURL,
		UserID:         rec.UserID,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
	}
}

func fromDomain(domain *social_medias.Domain) *SocialMedias {
	return &SocialMedias{
		ID:             domain.ID,
		Name:           domain.Name,
		SocialMediaURL: domain.SocialMediaURL,
		UserID:         domain.UserID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}
