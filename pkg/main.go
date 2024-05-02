package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

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
	// Initialize background periodic tasks
	scheduler.Schedule()
	// Configure endpoints
	http.HandleFunc("/health", controller.GetGealthCheck)
	// Start HTTP server
	serverConfig := profiles.GetSettings().Server
	serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	err = http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatalf("Error initializing HTTP server:\n  %s", err.Error())
	}
}
