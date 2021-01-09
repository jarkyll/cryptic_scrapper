package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-boilerplate/service"
	"net/http"
)

func HandleRegister(context *gin.Context) {
	request, exists := context.Get("request")
	if !exists {
		// TODO add logger
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	service.Register()


	// do action aka call services

	// get results


	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func HandleLogin(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "login message"})
}

func HandleLogout(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "log out message"})
}

func HandleRefreshToken(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "token refresh message"})
}