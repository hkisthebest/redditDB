package main

import (
	"fmt"
	"net/http"
	"os"
	"server/api"
	"server/db"
	"server/task"
	"time"

	m "server/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
  fmt.Println("Server starting on port: ", os.Getenv("PORT"))
  time.Local = time.UTC
  db.Migrate()
  db.Connect()
  defer db.Pool.Close()
  task.Init()
  r := chi.NewRouter()

  r.Use(m.SetContentType)
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Use(middleware.Timeout(60 * time.Second))
  r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://localhost:5173"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300,
  }))

  fmt.Println("??")
	r.Mount("/api", api.DatapointRouter())

	http.ListenAndServe(os.Getenv("PORT"), r)
}
