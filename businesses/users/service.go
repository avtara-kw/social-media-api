package users

import (
	"fmt"
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

func (us *userService) Login(userDomain *Domain) (*Domain, error) {
	res, err := us.userRepository.GetByEmail(userDomain.Email)

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, businesses.ErrAccountNotFound
		}
		return nil, err
	}

	if !utils.ValidateHash(userDomain.Password, res.Password) {
		return nil, businesses.ErrInvalidCredential
	}
	res.Token = utils.GenerateToken(res.ID, res.Email)
	return res, nil
}

func (us *userService) Delete(ID string) error {
	if err := us.userRepository.DeleteByID(ID); err != nil {
		fmt.Println(err)
		return businesses.ErrInternalServer
	}

	return nil
}

func (us *userService) Update(id string, userDomain *Domain) (*Domain, error) {
	var err error

	res, err := us.userRepository.UpdateUsernameAndEmail(id, userDomain.Email, userDomain.Username)
	fmt.Println(err)
	if err != nil {
		if strings.Contains(err.Error(), "users_username_key") {
			return nil, businesses.ErrUsernameAccountDuplicate
		} else if strings.Contains(err.Error(), "users_email_key") {
			return nil, businesses.ErrEmailAccountDuplicate
		} else if strings.Contains(err.Error(), "no rows") {
			return nil, businesses.ErrAccountNotFound
		}
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}
