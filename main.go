// Package main for Go-Tamboon
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Pachara-H/go-tamboon/internal/adapter/csv"
	"github.com/Pachara-H/go-tamboon/internal/adapter/omise"
	"github.com/Pachara-H/go-tamboon/internal/cipher"
	"github.com/Pachara-H/go-tamboon/internal/configs"
	"github.com/Pachara-H/go-tamboon/internal/reporter"
	"github.com/Pachara-H/go-tamboon/internal/services"
	"github.com/Pachara-H/go-tamboon/internal/validator"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("performing donations...")

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

	// Check command line arguments
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <csv-file-path>\n", os.Args[0])
		os.Exit(1)
	}

	// Load configuration
	configLoader := configs.NewLoader()
	cfg, err := configLoader.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	cfg.CSVFilePath = os.Args[1]
	defer configLoader.ClearConfig(cfg)

	log.Printf("Configuration loaded successfully")
	log.Printf("Environment: %s", cfg.Environment)
	log.Printf("Target .CSV file: %s", cfg.CSVFilePath)

	// Initialize and run donation service
	cipherAgent := cipher.NewAgent()
	validatorAgent := validator.NewAgent()
	reporterAgent := reporter.NewAgent()
	csvParser := csv.NewParser()
	omiseClient, err := omise.NewClient(cfg.Omise.PublicKey, cfg.Omise.SecretKey)
	if err != nil {
		log.Fatal(err)
	}

	svcWorker := services.New(cfg, cipherAgent, validatorAgent, reporterAgent, omiseClient, csvParser)
	go func() {
		defer cancel()
		if err := svcWorker.ProcessDonations(ctx); err != nil {
			log.Fatalf("Failed to process donations: %v", err.Error())
		}
	}()

	// Wait for context cancellation (graceful shutdown)
	<-ctx.Done()
	log.Println("Application shutdown complete.")
}
