package logger

import (
	constants "Fuzlex/src/share/const"
	"log"
	"os"
)

var logger *log.Logger

func GetLogger() *log.Logger {
	if logger == nil {
		initLogger()
	}
	return logger
}

func initLogger() {
	if isExistLogFile() {
		deleteLogFile()
	}
	logFile, err := os.Create(constants.LOG_FILE)
	if err != nil {
		panic(err)
	}
	logger = log.New(logFile, "", log.LstdFlags)
}

func isExistLogFile() bool {
	_, err := os.Stat(constants.LOG_FILE)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func deleteLogFile() {
	err := os.Remove(constants.LOG_FILE)
	if err != nil {
		panic(err)
	}
}
