package main

import (
	"GOProg1/internal/api"
	"log"
)

func main() {
	log.Println("Hello world!")
	api.StartServer()
	log.Println("Application terminated")
}
