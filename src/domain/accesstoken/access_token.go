package accesstoken

import (
	"strings"
	"time"

	"github.com/emadghaffari/bookstore_oauth-api/src/utils/errors"
)

const expireTime = 24

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.ResError  {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == ""{
		return errors.HandlerBagRequest("invalid access token")
	}
	if at.ClientID <=0 {
		return errors.HandlerBagRequest("invalid client id request")
	}
	if at.UserID <=0 {
		return errors.HandlerBagRequest("invalid user id request")
	}
	if at.Expires <=0 {
		return errors.HandlerBagRequest("invalid Expires request")
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
