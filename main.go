package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alikan97/Go-GRPC/repository"
)

func main() {
	db, err := repository.InitDB()

	if err != nil {
		log.Fatalf("Unable to initialize data source: %v\n", err)
	}

	router, err := inject(db)

	if err != nil {
		log.Fatalf("Failed to inject data sources %v", err)
	}

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v...\n", srv.Addr)

	// Wait for a sig interrupt (kills channel)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.DB.Close(); err != nil {
		log.Fatalf("Error occured during shutdown of data source %v", err)
	}

	log.Printf("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v\n", err)
	}
	// Set up a connection to the server.

	// r, err := c.GetAllAsset(context.Background(), &emptypb.Empty{})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("####### get server Greeting response: %v", r.GetResponse())
}
