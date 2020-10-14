package accesstoken

import "time"

const expireTime = 24

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
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
