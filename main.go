package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	server := http.Server{
		Handler: router,
		Addr: ":"+port,
	}

	fmt.Printf("Server has been running on port %v", port)
	
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}