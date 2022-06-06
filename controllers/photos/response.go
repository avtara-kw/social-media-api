package photos

import (
	"github.com/avtara-kw/social-media-api/businesses/photos"
	"time"
)

type ResponsePostPhoto struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	UserID    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type ResponseUpdatePhoto struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	UserID    int        `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ResponsePhotos struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	UserID    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
}

func ResponsePhotoPostFromDomain(domain *photos.Domain) *ResponsePostPhoto {
	return &ResponsePostPhoto{
		ID:        domain.ID,
		Title:     domain.Title,
		Caption:   domain.Caption,
		PhotoURL:  domain.PhotoURL,
		UserID:    domain.UserID,
		CreatedAt: domain.CreatedAt,
	}
}

func ResponsePhotoUpdateFromDomain(domain *photos.Domain) *ResponseUpdatePhoto {
	return &ResponseUpdatePhoto{
		ID:        domain.ID,
		Title:     domain.Title,
		Caption:   domain.Caption,
		PhotoURL:  domain.PhotoURL,
		UserID:    domain.UserID,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ResponsePhotosFromDomain(domain []photos.Domain) []ResponsePhotos {
	var res []ResponsePhotos

	for _, val := range domain {
		temp := ResponsePhotos{
			ID:        val.ID,
			Title:     val.Title,
			Caption:   val.Caption,
			PhotoURL:  val.PhotoURL,
			UserID:    val.UserID,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			User: struct {
				Username string `json:"username"`
				Email    string `json:"email"`
			}{val.Users.Username, val.Users.Email},
		}

		res = append(res, temp)
	}

	return res
}
