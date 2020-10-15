package http

import (
	"net/http"

	"github.com/emadghaffari/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/emadghaffari/bookstore_uesrs-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// NewHandler func
func NewHandler(service accesstoken.Service) AccessToken {
	return &accessTokenHandler{
		service: service,
	}
}

// AccessToken interface
type AccessToken interface {
	GetByID(*gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

func (ac *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := ac.service.GetByID(c.Param("access_token"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context)  {
	var ac accesstoken.AccessToken

	if err := c.ShouldBindJSON(&ac); err != nil {
		resErr := errors.HandlerBagRequest("invalid JSON format")
		c.JSON(resErr.Status, resErr)
		return
	}

	if err := handler.service.Create(ac); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, ac)
}