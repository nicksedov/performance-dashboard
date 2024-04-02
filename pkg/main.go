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

	http.HandleFunc("/health", controller.GetGealthCheck)

	// Start http server
	serverConfig := profiles.GetSettings().Server
	serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	srvErr := http.ListenAndServe(serverAddress, nil)
	if srvErr != nil {
		log.Fatalf("Error initializing HTTP server:\n  %s", srvErr.Error())
	}
}
