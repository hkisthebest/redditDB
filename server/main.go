package main

import (
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
  "log"
)

func main() {
  log.SetOutput(os.Stdout)
  port := os.Getenv("PORT")
  log.Println("Server started on port: ", port)
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

	r.Mount("/api", api.DatapointRouter())

  http.ListenAndServe(":3000", r)
}
