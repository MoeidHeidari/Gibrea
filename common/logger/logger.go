package common

import (
	"log"
	"os"

	"github.com/xtgo/uuid"
)

type JobReport struct {
	JobId     uuid.UUID
	Message   int
	ProcessId int
	LogLevel  LogLevel
}

type TaskReprot struct {
	taskId   uuid.UUID
	jobId    uuid.UUID
	message  string
	LogLevel LogLevel
}

//...................................................................
type LogLevel string

const (
	Error   LogLevel = "error"
	Debug   LogLevel = "debug"
	Warning LogLevel = "warning"
	Info    LogLevel = "info"
)

type PrintDestination string

const (
	File    PrintDestination = "file"
	Console PrintDestination = "console"
)

type LoggerOptions struct {
	FileAddress      string
	LogLevels        []LogLevel
	PrintDestination PrintDestination
}

var file os.File
var LogOptions LoggerOptions

//========================================================================================================================
func InitLogger(options LoggerOptions) error {
	LogOptions = options
	if options.PrintDestination == File {
		f, err := os.OpenFile(options.FileAddress, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		file = *f
		log.SetOutput(&file)
		log.Print("Gibrea log file")
		return err
	} else {
		log.SetOutput(os.Stdout)
		log.Print("Gibrea logging system")
	}

	return nil
}

//========================================================================================================================
func Finish() error {
	err := file.Close()
	return err
}
