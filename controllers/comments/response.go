package comments

import (
	"github.com/avtara-kw/social-media-api/businesses/comments"
	"time"
)

type ResponsePostComment struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   int        `json:"photo_id"`
	UserID    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type ResponseUpdateComment struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   int        `json:"photo_id"`
	UserID    int        `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ResponseComments struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   int        `json:"photo_id"`
	UserID    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
	Photo struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoURL string `json:"photo_url"`
		UserID   int    `json:"user_id"`
	} `json:"photo"`
}

func ResponseCommentPostFromDomain(domain *comments.Domain) *ResponsePostComment {
	return &ResponsePostComment{
		ID:        domain.ID,
		Message:   domain.Message,
		PhotoID:   domain.PhotoID,
		UserID:    domain.UserID,
		CreatedAt: domain.CreatedAt,
	}
}

func ResponseCommentUpdateFromDomain(domain *comments.Domain) *ResponseUpdateComment {
	return &ResponseUpdateComment{
		ID:        domain.ID,
		Message:   domain.Message,
		PhotoID:   domain.PhotoID,
		UserID:    domain.UserID,
		UpdatedAt: domain.CreatedAt,
	}
}

func ResponseCommentsFromDomain(domain []comments.Domain) []ResponseComments {
	var res []ResponseComments

	for _, val := range domain {
		temp := ResponseComments{
			ID:        val.ID,
			Message:   val.Message,
			UserID:    val.UserID,
			PhotoID:   val.PhotoID,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			User: struct {
				ID       int    `json:"id"`
				Username string `json:"username"`
				Email    string `json:"email"`
			}{int(val.Users.ID), val.Users.Username, val.Users.Email},
			Photo: struct {
				ID       int    `json:"id"`
				Title    string `json:"title"`
				Caption  string `json:"caption"`
				PhotoURL string `json:"photo_url"`
				UserID   int    `json:"user_id"`
			}{int(val.Photo.ID), val.Photo.Title, val.Photo.Caption, val.Photo.PhotoURL, val.Photo.UserID},
		}

		res = append(res, temp)
	}

	return res
}
