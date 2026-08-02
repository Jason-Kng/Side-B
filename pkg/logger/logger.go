package logger

import (
	"log"
	"os"
)

var logFile *os.File

func Init() error {
	var err error
	logFile, err = os.OpenFile(
		"debug.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0o644,
	)
	if err != nil {
		return err
	}
	log.SetOutput(logFile)
	return nil
}

func Close() error {
	if logFile == nil {
		return nil
	}
	return logFile.Close()
}

func Info(msg string, args ...any) {
	log.Printf("[INFO] "+msg, args...)
}

func Error(msg string, args ...any) {
	log.Printf("[ERROR] "+msg, args...)
}
