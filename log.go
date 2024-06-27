package golog

import (
	"os"
)

// Debug log debug line
func Debug(v ...any) {
	if debugLog == nil {
		return
	}
	debugLog.Println(v...)
}

// Info log info line
func Info(v ...any) {
	if infoLog == nil {
		return
	}
	infoLog.Println(v...)
}

// Warning log warning line
func Warning(v ...any) {
	if warningLog == nil {
		return
	}
	warningLog.Println(v...)
}

// Error log error line
func Error(v ...any) {
	if errorLog == nil {
		return
	}
	errorLog.Println(v...)
}

// Fatal log fatal error and terminate
func Fatal(v ...any) {
	Error(v...)
	Error("Fatal error")
	Close()
	os.Exit(1)
}

// Debugf log debug line with format
func Debugf(format string, v ...any) {
	if debugLog == nil {
		return
	}
	debugLog.Printf(format, v...)
}

// Infof log info line with format
func Infof(format string, v ...any) {
	if infoLog == nil {
		return
	}
	infoLog.Printf(format, v...)
}

// Warningf log warning line with format
func Warningf(format string, v ...any) {
	if warningLog == nil {
		return
	}
	warningLog.Printf(format, v...)
}

// Errorf log error line with format
func Errorf(format string, v ...any) {
	if errorLog == nil {
		return
	}
	errorLog.Printf(format, v...)
}

// Fatalf log fatal error with format
func Fatalf(format string, v ...any) {
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
