package main

import (
	"bombardaketak/controllers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	print("\n=============================\n")
	err := godotenv.Load()
	if err != nil {
		print("Error loading .env file")
	}
	controllers.InitViews()
	//controllers.InitAPI()
	port := os.Getenv("PORT")
	println("Starting server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	print("=============================\n")
}
