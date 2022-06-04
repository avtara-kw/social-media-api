package businesses

import "errors"

var (
	ErrInternalServer           = errors.New("something gone wrong, contact administrator")
	ErrEmailAccountDuplicate    = errors.New("email is already taken")
	ErrInvalidCredential        = errors.New("email or password does not match")
	ErrUsernameAccountDuplicate = errors.New("username is already taken")
)
