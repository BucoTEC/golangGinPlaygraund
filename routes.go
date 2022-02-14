package main

import "github.com/gin-gonic/gin"

func HelloWorld(c *gin.Context) {

	c.JSON(200, "Hello welcome to my web server using fresh")
}

func HelloByName(c *gin.Context) {
	name := c.Param("name")
	res := SayHello(name)

	c.JSON(200, res )
}