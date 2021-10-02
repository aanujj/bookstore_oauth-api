package app

import (
	"github.com/bookstore_oauth-api/src/client/cassandra"
	"github.com/bookstore_oauth-api/src/controller"
	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = *gin.Default()
)

func StartApp() {
	session, Dberr := cassandra.GetSession()
	if Dberr != nil {
		panic(Dberr)
	}
	session.Close()
	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := controller.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.CreateAT)

	router.Run(":8080")
}
