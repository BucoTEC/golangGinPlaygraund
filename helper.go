package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//parsing name variable into greet
func SayHello(name string) string {

	helloThere := fmt.Sprintf("Hello %v, nice to meet you!", name)

	return helloThere
}

//env fetcher
func GoDotEnvVariable(key string) string {

	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }