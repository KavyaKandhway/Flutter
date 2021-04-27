package main

import (
	"fmt"
	"log"
	"os"

	"furrble.com/backend/routers"
	"github.com/joho/godotenv"
)

//init gets called before the main function
func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

//Execution starts from main function
func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	r := routers.SetupRouter()
	port := os.Getenv("port")

	// For run on requested port
	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "9990" //localhost
	}
	type Job interface {
		Run()
	}

	r.Run(":" + port)

}
