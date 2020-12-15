package main

import (
	"VInstaller/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.PUT("/runcommand", api.Run)
	router.Run()
}
