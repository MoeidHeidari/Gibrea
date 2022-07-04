package common

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/google/uuid"
)

type Report struct {
	JobId     uuid.UUID   `json:"jobId"`
	Message   string      `json:"message"`
	ProcessId int         `json:"processId"`
	Level     LogLevel    `json:"logLevel"`
	Detail    interface{} `json:"detail"`
	TaskId    uuid.UUID   `json:"taskId"`
}

type MigrationReport struct {
	taskId   uuid.UUID
	RepoName string
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
		f, err := os.OpenFile(options.FileAddress, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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
func PrintLog(report Report) error {
	if contains(LogOptions.LogLevels, report.Level) {
		logInfo, _ := json.Marshal(report)
		log.Println(string(logInfo))
		return nil
	}
	return errors.New("log level is not allowed")
}

//========================================================================================================================
func contains(slice []LogLevel, item LogLevel) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[string(s)] = struct{}{}
	}

	_, ok := set[string(item)]
	return ok
}

//========================================================================================================================
func Finish() error {
	err := file.Close()
	return err
}
