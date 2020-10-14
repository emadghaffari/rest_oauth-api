package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCheckAccessTokenExpiredTime(t *testing.T) {
	assert.EqualValues(t, 24, expireTime, "expire time should be 24 hourse")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.False(t, at.isExpired(), "access token is expired")

	assert.EqualValues(t, "", at.AccessToken, "the new access token should not have deiend access token id")

	assert.True(t, at.UserID == 0, "the new access token should not have deiend UserID id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.isExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(time.Hour * 3).Unix()
	assert.False(t, at.isExpired(), "access token expireing 3 hourse after now")

}
