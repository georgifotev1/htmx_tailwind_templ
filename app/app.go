package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/georgifotev1/wms/internal/database"
)

type App struct {
	db     database.Queries
	router http.Handler
}

func New(conn *sql.DB) *App {
	app := &App{
		db: *database.New(conn),
	}
	app.newRouter()

	return app
}

func (a *App) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv("PORT")),
		Handler: a.router,
	}

	go func() {
		fmt.Println("Starting server")

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %s\n", err)
		}
	}()

	<-ctx.Done()

	fmt.Println("\nShutting down HTTP server gracefully...")
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		fmt.Printf("HTTP server shutdown error: %s\n", err)
	}

	fmt.Println("HTTP server stopped.")
}

func DBConnect() (*sql.DB, error) {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		return nil, fmt.Errorf("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("can't connect to database: %v", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("connection not established: %v", err)
	}

	fmt.Println("Connected to database")

	return conn, nil
}
