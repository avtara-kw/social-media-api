package photos

import "github.com/avtara-kw/social-media-api/businesses/photos"

type RequestPhotosPostAndUpdate struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required"`
}

func (rec *RequestPhotosPostAndUpdate) ToDomain(ID int) *photos.Domain {
	return &photos.Domain{
		Title:    rec.Title,
		Caption:  rec.Caption,
		PhotoURL: rec.PhotoURL,
		UserID:   ID,
	}
}
