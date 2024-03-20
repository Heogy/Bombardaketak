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
	print("Starting server at port 9980\n")
	log.Fatal(http.ListenAndServe(":9980", nil))
	print("=============================\n")
}
