// Package main for Go-Tamboon
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Pachara-H/go-tamboon/internal/configs"
	"github.com/joho/godotenv"
)

func main() {
	// Create context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Received shutdown signal, gracefully shutting down...")
		cancel()
	}()

	// Setup environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	// Load configuration
	configLoader := configs.NewLoader()
	cfg, err := configLoader.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Check command line arguments
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <csv-file-path>\n", os.Args[0])
		os.Exit(1)
	}

	cfg.CSVFilePath = os.Args[1]

	// Validate file exists
	if _, err := os.Stat(cfg.CSVFilePath); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s", cfg.CSVFilePath)
	}

	log.Printf("Configuration loaded successfully")
	log.Printf("Environment: %s", cfg.Environment)
	log.Printf("Target .CSV file: %s", cfg.CSVFilePath)
	log.Printf("Omise Base URL: %s", cfg.Omise.BaseURL)

	// TODO: Initialize and run donation service
	// This will be implemented in subsequent tasks

	log.Println("Application setup complete. Ready for donation processing implementation.")

	// Wait for context cancellation (graceful shutdown)
	<-ctx.Done()
	log.Println("Application shutdown complete.")
}
