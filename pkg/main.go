package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"performance-dashboard/pkg/controller"
	"performance-dashboard/pkg/database"
	"performance-dashboard/pkg/profiles"
	"performance-dashboard/pkg/scheduler"
	"performance-dashboard/pkg/util"
)

func main() {
	// Process command line options
	flag.Parse()
	// Initialize logging
	util.InitLog()
	//Initialize database
	err := database.InitializeDB()
	if err != nil {
		log.Fatalf("Error initializing database connection:\n  %s", err.Error())
	}
	defer database.CloseDB() 
	// Initialize background periodic tasks
	scheduler.Schedule()
	// Configure endpoints
	http.HandleFunc("/health", controller.GetHealthCheck)
	// Start HTTP server
	serverConfig := profiles.GetSettings().Server
	serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	server := &http.Server{Addr: serverAddress}
	go handleShutdown(server)
	if err = server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error running HTTP server:\n  %v", err)
	}
	log.Println("HTTP server stopped")
}

func handleShutdown(server *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP server shutdown error: %v", err)
	}
}
