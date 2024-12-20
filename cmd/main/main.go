package main

import (
	"fmt"
	"jimdel/pkg/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	port := ":" + os.Getenv("APP_PORT")
	fmt.Println("Server running on port", port)
	err := server.Run(port)
	if err != nil {
		panic(err)
	}
}
