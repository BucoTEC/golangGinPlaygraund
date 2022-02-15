package main

import (
	"net/http"

	"ginTEST/models"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {

	c.JSON(200, "Hello welcome to my web server using fresh on port: 5000")
}

func HelloByName(c *gin.Context) {
	name := c.Param("name")
	res := SayHello(name)

	c.JSON(200, res )
}


func CreateUser(c *gin.Context){
	var newUser models.User
	c.Bind(&newUser)
	c.JSON(http.StatusOK, newUser)
}