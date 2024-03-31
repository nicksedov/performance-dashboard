package util

import (
	"io"
	"log"
	"os"
	"strings"

	"performance-dashboard/pkg/profiles"
)

func InitLog() {
	settings := profiles.GetSettings()
	if strings.TrimSpace(settings.Logger.Filename) != "" {
		lumberjackLogger := &settings.Logger
		multiWriter := io.MultiWriter(os.Stderr, lumberjackLogger)
		log.SetFlags(log.LstdFlags)
		log.SetOutput(multiWriter)
	}
}
