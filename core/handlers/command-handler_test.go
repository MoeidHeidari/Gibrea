package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartHandlerFunction(t *testing.T) {
	config, err := StartHandler("../config.example.yaml", true)
	assert.NoError(t, err)
	assert.Equal(t, config, true)

	_, error := StartHandler("../config.yaml", true)
	assert.Nil(t, error)

}

func TestStopHandlerDunction(t *testing.T) {}

func TestPauseHandlerFunction(t *testing.T) {}

func TestResumHandlerFunction(t *testing.T) {}

func TestStatusHandlerFunction(t *testing.T) {}

func TestLogsHandlerFunction(t *testing.T) {}
