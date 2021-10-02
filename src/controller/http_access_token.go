package controller

import (
	"net/http"
	"strings"

	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	CreateAT(c *gin.Context)
	UpdateExpirationTime(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}

}

func (acth *accessTokenHandler) GetById(c *gin.Context) {
	//query param to catch token id from postman url or query
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := acth.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)

}

func (acth *accessTokenHandler) CreateAT(c *gin.Context) {
	var accessToken access_token.AccessToken

	if err := c.ShouldBindJSON(&accessToken); err != nil {
		restErr := errors.StatusBadRequestError("invalid json body ")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := acth.service.CreateAT(accessToken); err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusCreated, accessToken)

}

func (acth *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	//query param to catch token id from postman url or query
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := acth.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, accessToken)

}
