package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/muf002/go-basic/src/pkg/config"
	"github.com/muf002/go-basic/src/pkg/routes"
)

func main() {
	config.Connect()
	routes := routes.GetRoutes()
	http.Handle("/", routes)
	fmt.Printf("Starting server at port 9010\n")
	if err := http.ListenAndServe(":9010", nil); err != nil {
		log.Fatal(err)
	}
}
