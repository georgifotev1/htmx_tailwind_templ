package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/georgifotev1/wms/app"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := app.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("failed to close database:\n", err)
		}
	}()

	app := app.New(conn)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)

	go app.Start(ctx, &wg)

	signalCh := make(chan os.Signal, 1)

	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh

	cancel()

	wg.Wait()

	fmt.Println("Shutdown complete.")
}
