package db

import (
	"github.com/emadghaffari/rest_oauth-api/src/clients/cassandra"
	"github.com/emadghaffari/rest_oauth-api/src/domain/accesstoken"
	"github.com/emadghaffari/rest_oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

var (
	getAccessTokenQuery      = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	createAccesstokenQuery   = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?)"
	updateAccesstokenExpires = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

// NewRepository func, for create a new DB Repository
func NewRepository() Repository {
	return &repository{}
}

// Repository interface
type Repository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.ResError)
	Create(accesstoken.AccessToken) *errors.ResError
	Update(accesstoken.AccessToken) *errors.ResError
}

type repository struct{}

func (r repository) GetByID(id string) (*accesstoken.AccessToken, *errors.ResError) {

	var result accesstoken.AccessToken
	if err := cassandra.GetSesstion().Query(getAccessTokenQuery, id).Scan(&result.AccessToken, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.HandlerNotFoundError("access_token with this id not found")
		}
		return nil, errors.HandlerNotFoundError(err.Error())
	}

	return &result, nil
}

func (r repository) Create(ac accesstoken.AccessToken) *errors.ResError {

	if err := cassandra.GetSesstion().Query(createAccesstokenQuery, ac.AccessToken, ac.UserID, ac.ClientID, ac.Expires).Exec(); err != nil {
		return errors.HandlerInternalServerError("error in insert new accessToken to server")
	}
	return nil
}

func (r repository) Update(ac accesstoken.AccessToken) *errors.ResError {

	if err := cassandra.GetSesstion().Query(updateAccesstokenExpires, ac.Expires, ac.AccessToken).Exec(); err != nil {
		return errors.HandlerInternalServerError("error in insert new accessToken to server")
	}
	return nil
}
