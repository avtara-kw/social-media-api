package photos

import (
	"fmt"
	"github.com/avtara-kw/social-media-api/businesses"
	"strings"
)

type photoService struct {
	photoRepository Repository
}

func NewPhotosService(rep Repository) Service {
	return &photoService{
		photoRepository: rep,
	}
}

func (us *photoService) Post(photoDomain *Domain) (*Domain, error) {
	res, err := us.photoRepository.Store(photoDomain)
	if err != nil {
		fmt.Println(err)
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *photoService) GetAll() ([]Domain, error) {
	res, err := us.photoRepository.GetAll()

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, businesses.ErrDataNotFound
		}
		return nil, err
	}

	return res, nil
}

func (us *photoService) Delete(photoID, userID int) error {
	res, err := us.photoRepository.GetByID(userID)
	if err != nil {
		return businesses.ErrInternalServer
	}

	if res.UserID != photoID {
		return businesses.ErrForbiddenAccess
	}

	if err := us.photoRepository.DeleteByID(photoID); err != nil {
		return businesses.ErrInternalServer
	}

	return nil
}

func (us *photoService) Update(data *Domain, photoID string) (*Domain, error) {
	var err error

	res, err := us.photoRepository.Update(data, photoID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, businesses.ErrForbiddenAccess
		}
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}
