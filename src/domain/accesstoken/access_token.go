package accesstoken

import (
	"strings"
	"time"

	"github.com/emadghaffari/res_errors/errors"
)

const expireTime = 24

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// Validate method
func (at *AccessToken) Validate() errors.ResError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.HandlerBadRequest("invalid access token")
	}
	if at.ClientID <= 0 {
		return errors.HandlerBadRequest("invalid client id request")
	}
	if at.UserID <= 0 {
		return errors.HandlerBadRequest("invalid user id request")
	}
	if at.Expires <= 0 {
		return errors.HandlerBadRequest("invalid Expires request")
	}
	return nil
}

// GetNewAccessToken func
func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expireTime * time.Hour).Unix(),
	}
}

func (at AccessToken) isExpired() bool {
	now := time.Now().UTC()
	expireDate := time.Unix(at.Expires, 0)
	return now.After(expireDate)
}
