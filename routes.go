package main

import "github.com/gin-gonic/gin"

func HelloWorld(c *gin.Context) {
	c.JSON(200, "Hello welcome to my web server using fresh")
}