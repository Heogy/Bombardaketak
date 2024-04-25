package main

import (
	"bombardaketak/controllers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	print("\n=============================\n")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	controllers.InitViews()
	controllers.InitAPI()
	print("Starting server at port 9980\n")
	log.Fatal(http.ListenAndServe(":9980", nil))
	print("=============================\n")
}
