package controllers

import (
	"github.com/Hermes-chat-App/hermes-auth-server/internal/application"
	"github.com/gin-gonic/gin"
)

func VerifyCode(c *gin.Context) {
	var requestedCode application.VerifyCodeRequest

	if !ValidateRequest(c, &requestedCode) {
		return
	}

	valid, err := application.VerifyCode(&requestedCode)

	if err != nil {
		c.Errors = append(c.Errors, c.Error(err))
		return
	}

	c.JSON(200, valid)
}

func Login(c *gin.Context) {
	var loginRequest application.LoginRequest

	if !ValidateRequest(c, &loginRequest) {
		return
	}

	response, err := application.Login(&loginRequest)

	if err != nil {
		c.Errors = append(c.Errors, c.Error(err))
		return
	}

	c.JSON(200, response)
}
