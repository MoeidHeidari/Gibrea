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
func TestCreateNewTaskWithABunchOfMigrations(t *testing.T) {
	config, error := config.Load("../../config/fixtures/fixture-config.yaml")
	assert.NoError(t, error, "Parse example config failed.")
	tasks = nil

	for i := 0; i < 60023; i++ {
		var migration core.RepositoryMigration
		migration.Name = "test_name"
		migration.Link = "test_link"
		migration.Description = "test_description"
		migration.Stars = 20

		core.AddMigration(migration)
	}

	for _, page := range core.Pages {
		newTask(page, config.Config.Tasks[0].Task)
	}
	assert.Equal(t, len(tasks), 61)

}
