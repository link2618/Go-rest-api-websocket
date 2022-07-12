package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/link2618/go/rest-ws/handlers"
	"github.com/link2618/go/rest-ws/middleware"
	"github.com/link2618/go/rest-ws/server"
)

// Definir endpoints
func BindRoutes(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/posts/{postId}", handlers.GetPostByIDHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/posts/{postId}", handlers.UpdatePostByIdHandler(s)).Methods(http.MethodPut)
	r.HandleFunc("/posts/{postId}", handlers.DeletePostByIdHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)
}

func main() {
	// Cargar variables de entorno
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}

	s.Start(BindRoutes)
}
