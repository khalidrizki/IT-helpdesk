package main

import (
	"context"
	"errors"
	"net"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/khalidrizki/IT-helpdesk/back-end/accounts-service/internal/account"
)

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

	accountRepository := account.NewRepository()
}
