package core

import (
	"main/config"
	core "main/core/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewTask(t *testing.T) {
	config, error := config.Load("../../config/fixtures/fixture-config.yaml")
	assert.NoError(t, error, "Parse example config failed.")

	var page core.RepositoryMigrationPage
	var UID = newTask(page, config.Config.Tasks[0].Task)
	assert.NotZero(t, UID)

}
