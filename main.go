package main

import (
	apiHandlers "aiplus_golang/internal/adapters/api"
	repositories "aiplus_golang/internal/adapters/repositories"
	"aiplus_golang/internal/core/domain"
	service "aiplus_golang/internal/core/services"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	config, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	dsn := getEnv(config.Database.Key, config.Database.Url)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer pool.Close()

	repo := repositories.NewEmployeeRepository(pool)
	empService := service.NewEmployeeService(repo)
	mux := apiHandlers.AddHanledrs(empService)

	server := &http.Server{
		Addr:    ":" + fmt.Sprint(config.Server.Port),
		Handler: mux,
	}

	go func() {
		log.Println("Server starting on :"+config.Server.Host, config.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(shutdownCtx)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LoadConfig(path string) (*domain.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg domain.Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
