package core

import (
	"main/config"
	repository "main/core/repository"
	task "main/core/task"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

//########################################################################################################################
func TestAddNewJob(t *testing.T) {
	config, error := config.Load("../../config/fixtures/fixture-config.yaml")
	assert.NoError(t, error, "Parse example config failed.")
	task.Tasks = nil

	for i := 0; i < 60023; i++ {
		var migration repository.RepositoryMigration
		migration.Name = faker.Name()
		migration.Link = faker.URL()
		migration.Description = faker.Sentence()
		start5, _ := faker.RandomInt(5)
		migration.Stars = start5[0]

		repository.AddMigration(migration)
	}

	for _, page := range repository.Pages {
		task.NewTask(page, config.Config.Tasks[0].Task)
	}

	NewJob(task.Tasks, func(job Job) JobStatus {
		job.status = running
		// network call...
		return job.status
	})

	assert.Equal(t, len(Jobs), 1)
}
