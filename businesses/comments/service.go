package comments

import (
	"github.com/avtara-kw/social-media-api/businesses"
	"strings"
)

type commentsService struct {
	commentRepository Repository
}

func NewCommentsService(rep Repository) Service {
	return &commentsService{
		commentRepository: rep,
	}
}

func (us *commentsService) Post(photoDomain *Domain) (*Domain, error) {
	res, err := us.commentRepository.Store(photoDomain)
	if err != nil {
		if strings.Contains(err.Error(), "fk_comment_photo") {
			return nil, businesses.ErrPhotoNotFound
		}
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *commentsService) GetAll() ([]Domain, error) {
	res, err := us.commentRepository.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, businesses.ErrDataNotFound
		}
		return nil, err
	}

	return res, nil
}

func (us *commentsService) Delete(commentID, userID int) error {
	res, err := us.commentRepository.GetByID(userID)
	if err != nil {
		return businesses.ErrInternalServer
	}

	if res.UserID != userID {
		return businesses.ErrForbiddenAccess
	}
	if err := us.commentRepository.DeleteByID(commentID); err != nil {
		return businesses.ErrInternalServer
	}

	return nil
}

func (us *commentsService) Update(data *Domain, commentID string) (*Domain, error) {
	var err error

	res, err := us.commentRepository.Update(data, commentID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, businesses.ErrForbiddenAccess
		}
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}
