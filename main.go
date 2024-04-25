package main

import (
	"bombardaketak/controllers"
	"log"
	"net/http"
)

func main() {
	print("\n=============================\n")
	controllers.InitViews()
	controllers.InitAPI()
	print("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
	print("=============================\n")
}
