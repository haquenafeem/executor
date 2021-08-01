package api

import (
	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	IsSuccess bool `json:"is_success"`
}

func pingHandler(c *gin.Context) {
	c.JSON(200, PingResponse{
		IsSuccess: true,
	})
}
