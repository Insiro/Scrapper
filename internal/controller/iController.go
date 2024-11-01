package controller

import "github.com/gin-gonic/gin"

type IController interface {
	Init(parent *gin.RouterGroup) IController
}
