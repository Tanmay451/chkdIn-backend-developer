package main

import (
	"chkdIn-backend-developer/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("main: Failed to load .env file with error: ", err)
	}
	r := routes.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
