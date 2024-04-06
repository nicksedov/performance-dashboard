package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"performance-dashboard/pkg/profiles"
)

func InitLog() {
	settings := profiles.GetSettings()
	consoleMode := settings.Logger.Console.Mode
	fileLogger := &settings.Logger.File
	var logWriter io.Writer
	if strings.TrimSpace(fileLogger.Filename) != "" {
		if consoleMode == "stderr" {
			logWriter = io.MultiWriter(os.Stderr, fileLogger)
		} else if consoleMode == "stdout" {
			logWriter = io.MultiWriter(os.Stdout, fileLogger)
		} else {
			logWriter = io.Writer(fileLogger)
		}
	} else {
		if consoleMode == "stderr" {
			logWriter = io.Writer(os.Stderr)
		} else if consoleMode == "stdout" {
			logWriter = io.Writer(os.Stdout)
		} else {
			logWriter = io.Writer(io.Discard)
			fmt.Print("Warning: application logging is suppressed!")
		}
	}
	log.SetFlags(log.LstdFlags)
	log.SetOutput(logWriter)
}
