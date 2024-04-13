package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NazarTymiv/time-series-visualization/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	// cors-configuration
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	// routers
	router.Get("/data", handlers.GetDataHandler)

	

	fmt.Printf("Server has been running on port %v", port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		log.Fatal(err)
	}

}