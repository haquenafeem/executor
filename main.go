package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haquenafeem/executor/api"
	"github.com/haquenafeem/executor/platform"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	api.Run(router)

	platform := platform.Get()
	fmt.Println("Host :", platform.Host)
	fmt.Println("Operating System :", platform.OS)
	fmt.Println("Architecture :", platform.Arch)
	fmt.Println("Running On Port :", "3000")

	router.Run(":3000")
}
