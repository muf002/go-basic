package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/muf002/go-basic/pkg/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	routes := routes.GetRoutes()
	http.Handle("/", routes)
	fmt.Printf("Starting server at port 9010\n")
	if err := http.ListenAndServe(":9010", nil); err != nil {
		log.Fatal(err)
	}
}
