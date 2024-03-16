package main

import (
	"net/http"
	"os"
	"server/api"
	"server/db"
	"server/task"
	"time"

	m "server/middleware"

	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
  log.SetOutput(os.Stdout)
  time.Local = time.UTC
  db.Migrate()
  db.Connect()
  defer db.Pool.Close()
  task.Init()
  r := chi.NewRouter()

  r.Use(m.SetContentType)
  r.Use(m.MeasureResponseDuration)
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Use(middleware.Timeout(60 * time.Second))
  r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"http://*", "https://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300,
  }))

  r.Get("/metrics", promhttp.Handler().ServeHTTP)

  r.Mount("/api", api.DatapointRouter())

  port := os.Getenv("PORT")
  http.ListenAndServe(port, r)
  log.Println("Server started on port: ", port)
}
