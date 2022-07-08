package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitializeLogging() {
	logFile := os.Getenv("LOGGER_OUTPUT")
	logLevel := os.Getenv("LOGGER_LEVEL")

	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.WithFields(log.Fields{"method": "InitializeLogging()", "error": err.Error()}).
			Error("[Error] faild to open log file")
	}
	log.SetOutput(file)

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.WithFields(log.Fields{"method": "InitializeLogging()", "error": err.Error()}).
			Error("[Error] failed to set log level")
	}
	log.SetLevel(level)

	log.SetFormatter(&log.JSONFormatter{})

}
