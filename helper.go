package main

import "fmt"

func SayHello(name string) string {

	var helloThere string 
	helloThere= fmt.Printf("Hello %v, nice to meet you!", name)

	return helloThere
}