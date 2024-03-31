package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"performance-dashboard/pkg/controller"
	"performance-dashboard/pkg/profiles"
	"performance-dashboard/pkg/scheduler"
	"performance-dashboard/pkg/util"
)

func main() {
	
	flag.Parse()

	util.InitLog()
	
	scheduler.Schedule()

	// Start http server	
	config := profiles.GetSettings()
	serverAddress := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	http.HandleFunc("/health", controller.GetGealth)
	srvErr := http.ListenAndServe(serverAddress, nil)
	if srvErr != nil {
		log.Fatalf("Error initializing HTTP server:\n  %s", srvErr.Error())
	}
}
