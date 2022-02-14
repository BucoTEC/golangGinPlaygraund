package main

import "fmt"

func SayHello(name string) string {

	helloThere := fmt.Sprintf("Hello %v, nice to meet you!", name)

	return helloThere
}