package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haquenafeem/executor/api"
)

func main() {
	router := gin.New()
	api.Run(router)
	router.Run(":3000")
}
