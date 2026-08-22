package main

import (
	"brief-url/Internal/auth"
	"brief-url/configs"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := configs.NewConfig()

	mux := http.NewServeMux()
	auth.NewAuthHandler(mux, auth.AuthHandlerDeps{Config: config})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server is listening on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
