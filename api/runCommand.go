package api

import (
	command2 "VInstaller/pkg/command"
	"github.com/gin-gonic/gin"
)



func Run(c *gin.Context){
	c.Request.ParseMultipartForm(1000)
	ip := c.Request.PostForm["ip"][0]
	command := c.Request.PostForm["command"][0]
	output := command2.RunCommand(ip, command)
	c.Data(200, "text/html; charset=utf-8", []byte(output))
}


