package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/Rexkizzy22/evolve-credit/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// const version = "1.0.0"

type DB struct {
	dsn string
}

type config struct {
	port int
	env  string
	db   DB
}

type application struct {
	config config
	logger *log.Logger
	// models models.Models
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	// Initialize application variables
	cfg := config{
		port: 4000,
		env:  "development",
		db: DB{
			dsn: os.Getenv("DATABASE_URL"),
		},
	}

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModel(db),
	}

	// Create timeout settings for network requests to server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Server running on port", cfg.port)

	logger.Fatal(srv.ListenAndServe())
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
