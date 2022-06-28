package core

type RepositoryMigrationPage struct {
	page_repositories []RepositoryMigration
}

type RepositoryMigration struct {
	name        string
	link        string
	description string
	stars       int
}
