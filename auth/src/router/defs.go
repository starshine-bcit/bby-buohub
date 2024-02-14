package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starshine-bcit/bby-buohub/auth/service"
	"github.com/starshine-bcit/bby-buohub/auth/util"
	"gorm.io/gorm"
)

var Db *gorm.DB

func HandleLogin(c *gin.Context) {
	loginBody := new(util.AuthBody)
	if err := c.BindJSON(loginBody); err != nil {
		returnErr := util.ErrorResponse{
			ErrorName: "SerializationError",
			ErrorText: err.Error(),
		}
		c.JSON(http.StatusBadRequest, returnErr)
	}
	ok := service.ValidatePassword(Db, loginBody.Username, loginBody.Password)
	if !ok {
		returnErr := &util.ErrorResponse{
			ErrorName: "BadLogin",
			ErrorText: "User does not exist or wrong password",
		}
		c.JSON(http.StatusNotAcceptable, returnErr)
		return
	}
	username := loginBody.Username
	res := &util.LoginResponse{Valid: true}
	accessToken, err := service.NewAccessToken(
		service.GenUserAccessClaims(username))
	if err != nil {
		returnErr := &util.ErrorResponse{
			ErrorName: "TokenError",
			ErrorText: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, returnErr)
		return
	}
	refreshToken, err := service.NewRefreshToken(
		service.GenUserRefreshClaims(username))
	if err != nil {
		returnErr := &util.ErrorResponse{
			ErrorName: "TokenError",
			ErrorText: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, returnErr)
		return
	}
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	c.JSON(http.StatusAccepted, res)
}

func HandleCreate(c *gin.Context) {
	createBody := new(util.AuthBody)
	res := &util.CreateResponse{
		Created: false,
	}
	if err := c.BindJSON(createBody); err != nil {
		returnErr := &util.ErrorResponse{
			ErrorName: "SerializationError",
			ErrorText: err.Error(),
		}
		c.JSON(http.StatusBadRequest, returnErr)
		return
	}
	username := createBody.Username
	ok := service.CreateUser(Db, username, createBody.Password)
	if ok {
		res.Created = true
		accessToken, err := service.NewAccessToken(
			service.GenUserAccessClaims(username))
		if err != nil {
			returnErr := &util.ErrorResponse{
				ErrorName: "TokenError",
				ErrorText: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, returnErr)
			return
		}
		refreshToken, err := service.NewRefreshToken(
			service.GenUserRefreshClaims(username))
		if err != nil {
			returnErr := &util.ErrorResponse{
				ErrorName: "TokenError",
				ErrorText: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, returnErr)
			return
		}
		res.AccessToken = accessToken
		res.RefreshToken = refreshToken
		c.JSON(http.StatusCreated, res)
	} else {
		c.JSON(http.StatusBadRequest, res)
	}
}

func HandleAuth(c *gin.Context) {
	tokenBody := new(util.TokenBody)
	if err := c.BindJSON(tokenBody); err != nil {
		returnErr := util.ErrorResponse{
			ErrorName: "SerializationError",
			ErrorText: err.Error(),
		}
		c.JSON(http.StatusBadRequest, returnErr)
	}
	authResponse := &util.AuthResponse{Ok: true}
	status := http.StatusAccepted
	tokenBody, err := service.ValidateAuthRequest(tokenBody, Db)
	if err != nil {
		status = http.StatusNotAcceptable
		authResponse.Ok = false
	}
	if tokenBody != nil {
		authResponse.NewToken = tokenBody.AccessToken
		authResponse.Refreshed = true
	}
	c.JSON(status, authResponse)
}

func HandleRefresh(c *gin.Context) {
	tokenBody := new(util.TokenBody)
	if err := c.BindJSON(tokenBody); err != nil {
		returnErr := util.ErrorResponse{
			ErrorName: "SerializationError",
			ErrorText: err.Error(),
		}
		c.JSON(http.StatusBadRequest, returnErr)
	}
	authResponse := &util.AuthResponse{Ok: true}
	status := http.StatusAccepted
	tokenBody, err := service.ValidateAuthRequest(tokenBody, Db)
	if err != nil {
		status = http.StatusNotAcceptable
		authResponse.Ok = false
	}
	if tokenBody != nil {
		authResponse.NewToken = tokenBody.AccessToken
		authResponse.Refreshed = true
	}
	c.JSON(status, authResponse)
}
