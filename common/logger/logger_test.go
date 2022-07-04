package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogeerInitFunction(t *testing.T) {

	var options LoggerOptions
	options.PrintDestination = File
	options.FileAddress = "fixtures/fixturelog.log"

	err := InitLogger(options)

	assert.NoError(t, err)
}

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

	err := InitLogger(options)

	assert.NoError(t, err)
}

//========================================================================================================================
func TestEndLoggerFunction(t *testing.T) {
	err := Finish()
	assert.NoError(t, err)
}

//========================================================================================================================
