package users

import (
	"github.com/avtara-kw/social-media-api/businesses"
	"github.com/avtara-kw/social-media-api/utils"
	"strings"
)

type userService struct {
	userRepository Repository
}

func NewUserService(rep Repository) Service {
	return &userService{
		userRepository: rep,
	}
}

func (us *userService) Registration(userDomain *Domain) (*Domain, error) {
	var err error

	userDomain.Password, err = utils.HashPassword(userDomain.Password)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	res, err := us.userRepository.Store(userDomain)
	if err != nil {
		if strings.Contains(err.Error(), "users_username_key") {
			return nil, businesses.ErrUsernameAccountDuplicate
		} else if strings.Contains(err.Error(), "users_email_key") {
			return nil, businesses.ErrEmailAccountDuplicate
		}
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}
