package logger

import (
	"log"
	"os"
)

var (
	// Logger is the application-wide logger
	Logger *log.Logger
)

func init() {
	Logger = log.New(os.Stdout, "JobSchedulerService: ", log.LstdFlags|log.Lshortfile)
}

// Info logs informational messages
func Info(v ...interface{}) {
	Logger.SetPrefix("INFO: ")
	Logger.Println(v...)
}

// Error logs error messages
func Error(v ...interface{}) {
	Logger.SetPrefix("ERROR: ")
	Logger.Println(v...)
}

// Debug logs debug messages
func Debug(v ...interface{}) {
	Logger.SetPrefix("DEBUG: ")
	Logger.Println(v...)
}

// Fatal logs fatal messages
func Fatal(v ...interface{}) {
	Logger.SetPrefix("FATAL: ")
	Logger.Fatalln(v...)
}
