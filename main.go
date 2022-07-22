package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"evolve-credit/pkg/routes"
	"evolve-credit/pkg/utils"

	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type config struct {
	port int
	env  string
}

func main() {
	port, err := strconv.Atoi(utils.Getenv("PORT", "4000"))
	if err != nil {
		log.Fatal("Unable to parse port: ", err)
	}

	// Initialize application variables
	cfg := config{
		port: port,
		env:  "development",
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Create timeout settings for network requests to server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      routes.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Server running on port", cfg.port)

	logger.Fatal(srv.ListenAndServe())
}
