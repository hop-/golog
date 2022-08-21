package golog

import (
	"os"
)

// Debug log debug line
func Debug(v ...interface{}) {
	if debugLog == nil {
		return
	}
	debugLog.Println(v...)
}

// Info log info line
func Info(v ...interface{}) {
	if infoLog == nil {
		return
	}
	infoLog.Println(v...)
}

// Warning log warning line
func Warning(v ...interface{}) {
	if warningLog == nil {
		return
	}
	warningLog.Println(v...)
}

// Error log error line
func Error(v ...interface{}) {
	if errorLog == nil {
		return
	}
	errorLog.Println(v...)
}

// Fatal log fatal error and terminate
func Fatal(v ...interface{}) {
	Error(v...)
	Error("Fatal error")
	Close()
	os.Exit(1)
}

// Debugf log debug line with format
func Debugf(format string, v ...interface{}) {
	if debugLog == nil {
		return
	}
	debugLog.Printf(format, v...)
}

// Infof log info line with format
func Infof(format string, v ...interface{}) {
	if infoLog == nil {
		return
	}
	infoLog.Printf(format, v...)
}

// Warningf log warning line with format
func Warningf(format string, v ...interface{}) {
	if warningLog == nil {
		return
	}
	warningLog.Printf(format, v...)
}

// Errorf log error line with format
func Errorf(format string, v ...interface{}) {
	if errorLog == nil {
		return
	}
	errorLog.Printf(format, v...)
}

// Fatalf log fatal error with format
func Fatalf(format string, v ...interface{}) {
	Errorf(format, v...)
	Error("Fatal error")
	Close()
	os.Exit(1)
}

// Close close log file
func Close() {
	Info("Session closed.")

	if file != nil {
		file.Close()
	}
}
