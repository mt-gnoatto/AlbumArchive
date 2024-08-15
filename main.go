package main

import (
	"album-tracker/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := routes.SetupRoutes()

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
