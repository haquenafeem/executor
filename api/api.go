package api

import (
	"github.com/gin-gonic/gin"
)

func Run(router *gin.Engine) {
	router.GET("/ping", pingHandler)
	router.GET("/metadata", metaDataHandler)
	router.POST("/execute", executionHandler)
}
