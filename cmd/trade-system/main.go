package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"trade-organization/internal/application/management"
	"trade-organization/internal/infrastructure/database/repository"

	"trade-organization/internal/application/auth"
	"trade-organization/internal/application/reports"
	"trade-organization/internal/application/supply"
	"trade-organization/internal/application/trading"
	"trade-organization/internal/infrastructure/database"
	customhttp "trade-organization/internal/infrastructure/http"
	"trade-organization/internal/infrastructure/http/handlers"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	log.Println("Starting Trade System API...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/trade_db?sslmode=disable"
	}

	log.Println("Running database migrations...")
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	log.Println("Database migrated successfully!")

	dbPool, err := database.NewPostgresPool(ctx, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbPool.Close() // Закрываем пул при завершении работы сервера
	log.Println("Connected to PostgreSQL successfully")

	authRepo := repository.NewAuthRepository(dbPool)
	tradeRepo := repository.NewTradeRepository(dbPool)
	supplyRepo := repository.NewSupplyRepository(dbPool)
	reportRepo := repository.NewReportRepository(dbPool)
	storeRepo := repository.NewStoreRepository(dbPool)

	authService := auth.NewAuthService(authRepo)
	tradeService := trading.NewTradeService(tradeRepo)
	supplyService := supply.NewSupplyService(supplyRepo)
	reportService := reports.NewReportService(reportRepo)
	storeService := management.NewStoreService(storeRepo)

	authHandler := handlers.NewAuthHandler(authService)
	tradeHandler := handlers.NewTradeHandler(tradeService)
	supplyHandler := handlers.NewSupplyHandler(supplyService)
	reportHandler := handlers.NewReportHandler(reportService)
	storeHandler := handlers.NewStoreHandler(storeService)

	router := customhttp.SetupRouter(authHandler, tradeHandler, reportHandler, supplyHandler, storeHandler)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Println("Server is running on http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
