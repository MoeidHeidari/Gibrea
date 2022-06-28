package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMigration(t *testing.T) {
	for i := 0; i < 3500; i++ {
		var migration RepositoryMigration
		migration.Name = "test_name"
		migration.Link = "test_link"
		migration.Description = "test_description"
		migration.Stars = 20

		AddMigration(migration)
	}
	assert.Equal(t, len(Pages), 4)
	assert.Equal(t, len(Pages[0].page_repositories), 1000)
	assert.Equal(t, len(Pages[1].page_repositories), 1000)
	assert.Equal(t, len(Pages[2].page_repositories), 1000)
	assert.Equal(t, len(Pages[3].page_repositories), 500)
}

func TestMigrationWithZeroPages(t *testing.T) {
	Pages = nil
	var migration RepositoryMigration
	migration.Name = "test_name"
	migration.Link = "test_link"
	migration.Description = "test_description"
	migration.Stars = 20

	AddMigration(migration)
	assert.Equal(t, len(Pages), 1)
}
func TestAddbulkMigration(t *testing.T) {
	for i := 0; i < 60023; i++ {
		var migration RepositoryMigration
		migration.Name = "test_name"
		migration.Link = "test_link"
		migration.Description = "test_description"
		migration.Stars = 20

		AddMigration(migration)
	}
	assert.Equal(t, len(Pages), 61)
	assert.Equal(t, len(Pages[0].page_repositories), 1000)
	assert.Equal(t, len(Pages[1].page_repositories), 1000)
	assert.Equal(t, len(Pages[2].page_repositories), 1000)
	assert.Equal(t, len(Pages[60].page_repositories), 24)
}
