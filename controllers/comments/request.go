package comments

import (
	"github.com/avtara-kw/social-media-api/businesses/comments"
)

type RequestCommentPost struct {
	Message string `json:"message" binding:"required"`
	PhotoID int    `json:"photo_id" binding:"required"`
}

type RequestCommentUpdate struct {
	Message string `json:"message" binding:"required"`
}

func (rec *RequestCommentPost) ToDomain(ID int) *comments.Domain {
	return &comments.Domain{
		Message: rec.Message,
		PhotoID: rec.PhotoID,
		UserID:  ID,
	}
}

func (rec *RequestCommentUpdate) ToDomain(ID int) *comments.Domain {
	return &comments.Domain{
		Message: rec.Message,
		UserID:  ID,
	}
}
