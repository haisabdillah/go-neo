package logging

import (
	"log/slog"
	"os"
)

var (
	Logger *slog.Logger
)

func InitLog() error {
	// Create or open the log file
	// Define the log directory and file
	logDir := "storage/log"
	logFilePath := logDir + "/app.log"

	// Create the log directory if it doesn't exist
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return err
	}

	// Create or open the log file
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	Logger = slog.New(slog.NewJSONHandler(logFile, nil))
	return nil
}
