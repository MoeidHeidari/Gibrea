package core

import (
	"main/config"
	core "main/core/repository"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

//========================================================================================================================
func TestCreateNewTask(t *testing.T) {
	config, error := config.Load("../../config/fixtures/fixture-config.yaml")
	assert.NoError(t, error, "Parse example config failed.")

	var page core.RepositoryMigrationPage
	var UID = NewTask(page, config.Config.Tasks[0].Task)
	assert.NotZero(t, UID)

}

//========================================================================================================================
func TestCreateNewTaskWithABunchOfMigrations(t *testing.T) {
	config, error := config.Load("../../config/fixtures/fixture-config.yaml")
	assert.NoError(t, error, "Parse example config failed.")

	Tasks = nil

	for i := 0; i < 60023; i++ {
		var migration core.RepositoryMigration
		migration.Name = faker.Name()
		migration.Link = faker.URL()
		migration.Description = faker.Sentence()
		start5, _ := faker.RandomInt(5)
		migration.Stars = start5[0]

		core.AddMigration(migration)
	}

	for _, page := range core.Pages {
		NewTask(page, config.Config.Tasks[0].Task)
	}
	assert.Equal(t, len(Tasks), 61)
}
