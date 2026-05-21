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
	return logFile.Close()
}

func Info(msg string) {
	log.Println("[INFO]", msg)
}

func Error(msg string) {
	log.Println("[ERROR]", msg)
}
