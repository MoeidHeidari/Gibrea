package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMigration(t *testing.T) {
	for i := 0; i < 3500; i++ {
		var migration RepositoryMigration
		migration.name = "test_name"
		migration.link = "test_link"
		migration.description = "test_description"
		migration.stars = 20

		AddMigration(migration)
	}
	assert.Equal(t, len(pages), 4)
	assert.Equal(t, len(pages[0].page_repositories), 1000)
	assert.Equal(t, len(pages[1].page_repositories), 1000)
	assert.Equal(t, len(pages[2].page_repositories), 1000)
	assert.Equal(t, len(pages[3].page_repositories), 500)
}

func TestMigrationWithZeroPages(t *testing.T) {
	pages = nil
	var migration RepositoryMigration
	migration.name = "test_name"
	migration.link = "test_link"
	migration.description = "test_description"
	migration.stars = 20

	AddMigration(migration)
	assert.Equal(t, len(pages), 1)
}
func TestAddbulkMigration(t *testing.T) {
	for i := 0; i < 60023; i++ {
		var migration RepositoryMigration
		migration.name = "test_name"
		migration.link = "test_link"
		migration.description = "test_description"
		migration.stars = 20

		AddMigration(migration)
	}
	assert.Equal(t, len(pages), 61)
	assert.Equal(t, len(pages[0].page_repositories), 1000)
	assert.Equal(t, len(pages[1].page_repositories), 1000)
	assert.Equal(t, len(pages[2].page_repositories), 1000)
	assert.Equal(t, len(pages[60].page_repositories), 24)
}
