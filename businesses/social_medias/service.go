package social_medias

import (
	"fmt"
	"github.com/avtara-kw/social-media-api/businesses"
	"strings"
)

type socialMediaService struct {
	socialMediaRepository Repository
}

func NewSocialMediasService(rep Repository) Service {
	return &socialMediaService{
		socialMediaRepository: rep,
	}
}

func (us *socialMediaService) Post(data *Domain) (*Domain, error) {
	res, err := us.socialMediaRepository.Store(data)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *socialMediaService) GetAll(ID string) ([]Domain, error) {
	res, err := us.socialMediaRepository.GetAllByID(ID)

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, businesses.ErrDataNotFound
		}
		return nil, err
	}
	return res, nil
}

func (us *socialMediaService) Update(data *Domain, socialMediaID string) (*Domain, error) {
	var err error

	res, err := us.socialMediaRepository.Update(data, socialMediaID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, businesses.ErrForbiddenAccess
		}
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *socialMediaService) Delete(socialMediaID, userID int) error {
	res, err := us.socialMediaRepository.GetByID(userID)
	if err != nil {
		return businesses.ErrInternalServer
	}

	if res.UserID != userID {
		fmt.Println(res.UserID)
		return businesses.ErrForbiddenAccess
	}

	if err := us.socialMediaRepository.DeleteByID(socialMediaID); err != nil {
		return businesses.ErrInternalServer
	}

	return nil
}
