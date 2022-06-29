package core

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

//########################################################################################################################
func TestAddMigration(t *testing.T) {
	for i := 0; i < 3500; i++ {
		var migration RepositoryMigration
		migration.Name = faker.Name()
		migration.Link = faker.URL()
		migration.Description = faker.Sentence()
		start5, _ := faker.RandomInt(5)
		migration.Stars = start5[0]

		AddMigration(migration)
	}
	assert.Equal(t, len(Pages), 4)
	assert.Equal(t, len(Pages[0].page_repositories), 1000)
	assert.Equal(t, len(Pages[1].page_repositories), 1000)
	assert.Equal(t, len(Pages[2].page_repositories), 1000)
	assert.Equal(t, len(Pages[3].page_repositories), 500)
}

//========================================================================================================================
func TestMigrationWithZeroPages(t *testing.T) {
	Pages = nil
	var migration RepositoryMigration
	migration.Name = faker.Name()
	migration.Link = faker.URL()
	migration.Description = faker.Sentence()
	start5, _ := faker.RandomInt(5)
	migration.Stars = start5[0]

	AddMigration(migration)
	assert.Equal(t, len(Pages), 1)
}

//========================================================================================================================
func TestAddbulkMigration(t *testing.T) {
	for i := 0; i < 60023; i++ {
		var migration RepositoryMigration
		migration.Name = faker.Name()
		migration.Link = faker.URL()
		migration.Description = faker.Sentence()
		start, _ := faker.RandomInt(5)
		migration.Stars = start[0]

		AddMigration(migration)
	}
	assert.Equal(t, len(Pages), 61)
	assert.Equal(t, len(Pages[0].page_repositories), 1000)
	assert.Equal(t, len(Pages[1].page_repositories), 1000)
	assert.Equal(t, len(Pages[2].page_repositories), 1000)
	assert.Equal(t, len(Pages[60].page_repositories), 24)
}

//========================================================================================================================
func TestNewMigration(t *testing.T) {
	stars, _ := faker.RandomInt(1)
	var migration = NewMigration(faker.Name(), faker.URL(), faker.Sentence(), stars[0], func(migration RepositoryMigration) {
		// do any operation...
	})

	assert.NotNil(t, migration)
	assert.Equal(t, migration.Stars, stars[0])
}
