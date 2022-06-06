package social_medias

import (
	"github.com/avtara-kw/social-media-api/businesses/social_medias"
	"time"
)

type ResponseSocialMedia struct {
	ID             uint       `json:"id"`
	Name           string     `json:"title"`
	SocialMediaURL string     `json:"caption"`
	UserID         int        `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
}

type ResponseUpdateSocialMedia struct {
	ID             uint       `json:"id"`
	Name           string     `json:"title"`
	SocialMediaURL string     `json:"caption"`
	UserID         int        `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type ResponseSocialMedias struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaURL string     `json:"social_media_url"`
	UserID         int        `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	User           struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
}

func ResponseSocialMediasPostFromDomain(domain *social_medias.Domain) *ResponseSocialMedia {
	return &ResponseSocialMedia{
		ID:             domain.ID,
		Name:           domain.Name,
		SocialMediaURL: domain.SocialMediaURL,
		UserID:         domain.UserID,
		CreatedAt:      domain.CreatedAt,
	}
}

func ResponseSocialMediaUpdateFromDomain(domain *social_medias.Domain) *ResponseUpdateSocialMedia {
	return &ResponseUpdateSocialMedia{
		ID:             domain.ID,
		Name:           domain.Name,
		SocialMediaURL: domain.SocialMediaURL,
		UserID:         domain.UserID,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func ResponseSocialMediasFromDomain(domain []social_medias.Domain) []ResponseSocialMedias {
	var res []ResponseSocialMedias

	for _, val := range domain {
		temp := ResponseSocialMedias{
			ID:             val.ID,
			Name:           val.Name,
			SocialMediaURL: val.SocialMediaURL,
			UserID:         val.UserID,
			CreatedAt:      val.CreatedAt,
			UpdatedAt:      val.UpdatedAt,
			User: struct {
				ID       int    `json:"id"`
				Username string `json:"username"`
			}{val.UserID, val.Users.Username},
		}

		res = append(res, temp)
	}

	return res
}
