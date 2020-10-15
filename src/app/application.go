package app

import (
	"github.com/emadghaffari/bookstore_oauth-api/src/clients/cassandra"
	"github.com/emadghaffari/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/emadghaffari/bookstore_oauth-api/src/domain/repository/db"
	"github.com/emadghaffari/bookstore_oauth-api/src/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication func
func StartApplication() {
	sesstion, err := cassandra.GetSesstion()
	if err != nil {
		panic("cassandra connection lost")

	}
	defer sesstion.Close()
	atService := accesstoken.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token", atHandler.GetByID)
	router.POST("/oauth/access_token/", atHandler.Create)

	router.Run(":8080")
}