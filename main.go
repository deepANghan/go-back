package main

import (
	"fmt"
	"go-back/router"
	"log"
	"net/http"
)

func main() {

	server := &http.Server{
		Addr:    ":8080",
		Handler: router.SetUpRouter(),
	}

	fmt.Println("Go Server Starting on 8080")

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		log.Fatal(err)
	}
}
