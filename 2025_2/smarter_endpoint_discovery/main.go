// main.go
// Entry point for the Gorilla Mux + SQLite demo web app.
// Provides a small REST API with health, list items, and create item endpoints.
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"web_app/internal/db"
	"web_app/internal/handlers"
)

// jsonResponse is a helper to write JSON responses consistently.
func jsonResponse(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func main() {
	// Configuration via environment variables with sensible defaults.
	port := getenv("PORT", "8080")
	dbPath := getenv("DB_PATH", "data/app.db")

	// Initialize DB (creates file and table if needed).
	dbConn, err := db.OpenAndMigrate(dbPath)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer dbConn.Close()

	// Create router and wire handlers.
	r := mux.NewRouter()

	// Health endpoint
	r.HandleFunc("/health", handlers.HealthHandler).Methods(http.MethodGet)

	// Items endpoints
	itemHandler := handlers.NewItemHandler(dbConn)
	r.HandleFunc("/items", itemHandler.ListItems).Methods(http.MethodGet)
	r.HandleFunc("/items", itemHandler.CreateItem).Methods(http.MethodPost)

	// HTTP server with sensible timeouts.
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      loggingMiddleware(r),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown setup.
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("server listening on http://localhost:%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for termination signal
	<-shutdownCh
	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	} else {
		log.Println("server stopped")
	}
}

// loggingMiddleware logs basic request info.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// getenv returns the env var value or a default fallback.
func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
