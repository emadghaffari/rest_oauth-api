package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Start user repository tests.")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromAPI(t *testing.T) {
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := userRepository{}
	user, err := repository.Login("email@gmail.com", "the-password")
	fmt.Println(user, err)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid response when trying to login", err.Message)
}

func TestLoginInvalidErrorResponse(t *testing.T) {
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login credentials", "status": 404, "error":"not_found"}`,
	})

	repository := userRepository{}
	user, err := repository.Login("email@gmail.com", "the-password")
	fmt.Println(user, err)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface", err.Message)
}

func TestLoginInvalidCredentials(t *testing.T) {
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login credentials", "status": 404, "error":"not_found"}`,
	})

	repository := userRepository{}
	user, err := repository.Login("email@gmail.com", "the-password")
	fmt.Println(user, err)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid login Credentials", err.Message)
}

func TestLoginInvalidUnmarshalJsonResponse(t *testing.T) {
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":17,"email":"moha98ssss@gmail.com","created_at":"2020-10-12","status":"DEACTIVE"}`,
	})

	repository := userRepository{}
	user, err := repository.Login("email@gmail.com", "the-password")
	fmt.Println(user, err)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid login Unmarshal JSON Response", err.Message)
}

func TestLoginUserNotFoundResponse(t *testing.T) {
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":17,"email":"moha98ssss@gmail.com","first_name":"emad","status":"DEACTIVE"}`,
	})

	repository := userRepository{}
	user, err := repository.Login("email@gmail.com", "the-password")
	fmt.Println(user, err)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, 17, user.ID)
	assert.EqualValues(t, "moha98ssss@gmail.com", user.Email)
	assert.EqualValues(t, "email", user.FirstName)
}
