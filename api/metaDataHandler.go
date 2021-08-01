package api

import (
	"github.com/gin-gonic/gin"
	"github.com/haquenafeem/executor/platform"
)

func metaDataHandler(c *gin.Context) {
	c.JSON(200, platform.Get())
}
