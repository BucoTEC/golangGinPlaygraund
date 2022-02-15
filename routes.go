package main

import (
	"fmt"
	"net/http"

	"ginTEST/models"

	"github.com/gin-gonic/gin"
)




func HelloWorld(c *gin.Context) {
	port := GoDotEnvVariable("PORT")
	greet := fmt.Sprintf("Hello welcome to my web server using fresh on port: %v", port)
	c.JSON(200, greet  )
}

func HelloByName(c *gin.Context) {
	name := c.Param("name")
	res := SayHello(name)

	c.JSON(200, res )
}


func CreateUser(c *gin.Context){
	var newUser models.User
	c.Bind(&newUser)

	c.JSON(http.StatusOK, gin.H{
		"message":"newMesage",
		"NewUser": newUser,
	})
}