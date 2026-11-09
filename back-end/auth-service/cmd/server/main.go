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

	"github.com/khalidrizki/IT-helpdesk/back-end/auth-service/internal/account"
)

type server struct {
	accountService *account.Service
}

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

	log.Println("Successfully connected to PostgreSQL")

	accountRepository := account.NewRepository(dbPool)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return errors.New("JWT_SECRET environment variable is not set")
	}
	accountService := account.NewService(accountRepository, jwtSecret)

	s := &server{
		accountService: accountService,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/register", s.handleRegister)
	mux.HandleFunc("POST /auth/login", s.handleLogin)

	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Auth service running on http://localhost:8080")

	err = httpServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}
