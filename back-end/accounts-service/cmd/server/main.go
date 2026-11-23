package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/khalidrizki/IT-helpdesk/back-end/accounts-service/internal/account"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	if dbUser == "" || dbPassword == "" ||
		dbHost == "" || dbPort == "" || dbName == "" {
		return errors.New("database environment variables are incomplete")
	}

	dbURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(dbUser, dbPassword),
		Host:   net.JoinHostPort(dbHost, dbPort),
		Path:   dbName,
	}

	query := dbURL.Query()
	query.Set("sslmode", "disable")
	dbURL.RawQuery = query.Encode()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, dbURL.String())
	if err != nil {
		return err
	}
	defer dbPool.Close()

	if err := dbPool.Ping(ctx); err != nil {
		return err
	}

	log.Println("Successfully connected to PostgresSQL")

	accountRepository := account.NewRepository(dbPool)
	accountService := account.NewService(accountRepository)

	s := &server{
		accountService: accountService,
	}

	mux := http.NewServeMux()

	mux.HandleFunc(
		"POST /internal/accounts",
		s.handleCreateAccount,
	)

	httpServer := &http.Server{
		Addr:              ":8081",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Accounts service running on http://localhost:8081")

	err = httpServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}
