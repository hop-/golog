package golog

import (
	"io"
	"log"
	"os"
)

var (
	file       *os.File
	debugLog   *log.Logger
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
)

type LogMode = int

const (
	ErrorMode LogMode = iota + 1
	WarningMode
	InfoMode
	DebugMode
)

// Init initiate output streams
func Init(mode string) {
	switch getMode(mode) {
	case DebugMode:
		debugLog = log.New(os.Stdout, "Debug: ", log.Ldate|log.Ltime)
		fallthrough
	case InfoMode:
		infoLog = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime)
		fallthrough
	case WarningMode:
		warningLog = log.New(os.Stderr, "Warning: ", log.Ldate|log.Ltime)
		fallthrough
	default:
		errorLog = log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime)
	}

}

// InitFile initiate output streams and setup log file
func InitFile(fileName string, mode string) {
	var err error

	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	stdOutWriter := io.MultiWriter(file, os.Stdout)
	stdErrWriter := io.MultiWriter(file, os.Stderr)

	switch getMode(mode) {
	case DebugMode:
		debugLog = log.New(stdOutWriter, "Debug: ", log.Ldate|log.Ltime)
		fallthrough
	case InfoMode:
		infoLog = log.New(stdOutWriter, "Info: ", log.Ldate|log.Ltime)
		fallthrough
	case WarningMode:
		warningLog = log.New(stdErrWriter, "Warning: ", log.Ldate|log.Ltime)
		fallthrough
	default:
		errorLog = log.New(stdErrWriter, "Error: ", log.Ldate|log.Ltime)
	}
}

func getMode(mode string) LogMode {
	switch mode {
	case "WARNING":
		return WarningMode
	case "ERROR":
		return ErrorMode
	case "DEBUG":
		return DebugMode
	default:
		return InfoMode
	}
}
