package rest

import (
	"encoding/json"
	"time"

	"github.com/emadghaffari/bookstore_oauth-api/src/domain/users"
	"github.com/emadghaffari/bookstore_oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	restC = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

// UserRepository interface
type UserRepository interface {
	Login(string, string) (*users.User, *errors.ResError)
}

type userRepository struct{}

// NewRepository func, for create a new DB Repository
func NewRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Login(email string, password string) (user *users.User, resErr *errors.ResError) {
	request := users.LoginRequest{Email: email, Password: password}
	response := restC.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.HandlerInternalServerError("invalid resClient response when trying to send request")
	}
	if response.StatusCode > 299 {
		err := json.Unmarshal(response.Bytes(), &resErr)
		if err != nil {
			return nil, errors.HandlerInternalServerError("invalid error interface when trying to login")
		}
		return nil, resErr
	}
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.HandlerInternalServerError("invalid error trying to unmarshal user response")
	}

	return
}
