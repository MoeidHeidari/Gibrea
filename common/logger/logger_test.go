package common

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

//========================================================================================================================
func TestLoggerInitFunctionWithWrongFileAddress(t *testing.T) {
	var options LoggerOptions
	options.PrintDestination = File
	options.FileAddress = "fixtures_notavailable/fixturelog.log"

	err := InitLogger(options)

	assert.Error(t, err)
}

//========================================================================================================================
func TestLoggerInitFunctionWithoutFileAddress(t *testing.T) {
	var options LoggerOptions
	options.PrintDestination = Console
	options.FileAddress = "fixtures/fixturelog.log"
	options.LogLevels = append(options.LogLevels, Error)

	err := InitLogger(options)

	assert.NoError(t, err)
}

//========================================================================================================================
func TestLogeerInitFunction(t *testing.T) {

	var options LoggerOptions
	options.PrintDestination = File
	options.FileAddress = "fixtures/fixturelog.log"
	options.LogLevels = append(options.LogLevels, Error)
	options.LogLevels = append(options.LogLevels, Warning)

	err := InitLogger(options)

	assert.NoError(t, err)
}

//========================================================================================================================
func TestLogFunctionWithjobReportInterface(t *testing.T) {

	var report Report
	processId, _ := faker.RandomInt(2)
	report.JobId = uuid.New()
	report.Message = "test job logg message"
	report.ProcessId = processId[0]
	report.Level = Error
	logError := PrintLog(report)
	assert.NoError(t, logError)
}

//========================================================================================================================
func TestLogFunctionShouldThrowLogLevelisNotAllowed(t *testing.T) {

	var report Report
	processId, _ := faker.RandomInt(2)
	report.JobId = uuid.New()
	report.Message = "test job logg message"
	report.ProcessId = processId[0]
	report.Level = Warning
	logError := PrintLog(report)
	assert.NoError(t, logError)
}

//========================================================================================================================
func TestEndLoggerFunction(t *testing.T) {
	err := Finish()
	assert.NoError(t, err)
}
