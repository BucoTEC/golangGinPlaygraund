package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SayHello(name string) string {

	helloThere := fmt.Sprintf("Hello %v, nice to meet you!", name)

	return helloThere
}
func GoDotEnvVariable(key string) string {

	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }