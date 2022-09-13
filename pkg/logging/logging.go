package logging

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Logger() *log.Logger {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.DebugLevel)
	return logger
}
