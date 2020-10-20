package http

import (
	"net/http"

	"github.com/emadghaffari/res_errors/errors"
	"github.com/emadghaffari/rest_oauth-api/src/domain/accesstoken"
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

func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := handler.service.GetByID(c.Param("access_token"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}

// Create method
func (handler *accessTokenHandler) Create(c *gin.Context) {
	var ac accesstoken.AccessToken

	if err := c.ShouldBindJSON(&ac); err != nil {
		resErr := errors.HandlerBadRequest("invalid JSON format")
		c.JSON(resErr.Status(), resErr)
		return
	}

	if err := handler.service.Create(ac); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, ac)
}
