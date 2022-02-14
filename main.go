package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/", HelloWorld)
	server.GET("/:name", HelloByName)


	server.Run(":5000")
}