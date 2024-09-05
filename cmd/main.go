package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/saidazimqilichov/fitnas-tracking-app/storage"
	"github.com/sqlc-dev/pqtype"
)

func main() {
	connString := "postgres://postgres:7777@localhost:5432/fitnes_tracking_app?sslmode=disable"
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := sql.Open("postgres", connString)
	if err != nil {
		logger.Error("failed to connect")
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Error("failed to ping")
		os.Exit(1)
	}

	ctx := context.Background()
	queries := storage.New(db)

	err = queries.CreateUser(ctx, storage.CreateUserParams{
		Username: sql.NullString{String: "test"},
		Email: sql.NullString{String: "test@gmail.com"},
		Profile: pqtype.NullRawMessage{},
	})
	if err != nil {
		logger.Error("failed to create user")
		os.Exit(1)
	}

	users, err := queries.ListUsers(ctx)
	if err != nil {
		logger.Error("failed to list users")
		os.Exit(1)
	}

	fmt.Println("users: ", users)

}
