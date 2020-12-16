package main

import (
	"VInstaller/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/runcommand", api.Run)
	router.Run()
}
