package accesstoken

import "github.com/emadghaffari/bookstore_oauth-api/src/utils/errors"

// Service interface
type Service interface {
	GetByID(string) (*AccessToken, errors.ResError)
}
