package social_medias

import (
	"github.com/avtara-kw/social-media-api/businesses/social_medias"
)

type RequestSocialMediasPostAndUpdate struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" bindung:"required"`
}

func (rec *RequestSocialMediasPostAndUpdate) ToDomain(ID int) *social_medias.Domain {
	return &social_medias.Domain{
		Name:           rec.Name,
		SocialMediaURL: rec.SocialMediaURL,
		UserID:         ID,
	}
}
