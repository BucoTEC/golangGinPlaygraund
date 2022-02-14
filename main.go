package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/", HelloWorld)

	server.Run(":5000")
}