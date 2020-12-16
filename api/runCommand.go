package api

import (
	"VInstaller/pkg/command"
	"fmt"
	"github.com/gin-gonic/gin"
)

// The PUT request function that runs the command
func Run(c *gin.Context){
	var jsonData map[string]string
	c.BindJSON(&jsonData)
	fmt.Println(jsonData)
	output := command.RunCommand(jsonData["ip"], jsonData["command"])
	c.JSON(200, output)

}


