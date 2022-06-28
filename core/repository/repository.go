package core

import (
	"main/config"
)

type StartMigration func(options config.Config)
type CancelMigration func()

type RepositoryMigrationPage struct {
	page_repositories []RepositoryMigration
}

//...................................................................
type RepositoryMigration struct {
	Name            string
	Link            string
	Description     string
	Stars           int
	StartMigration  StartMigration
	CancelMigration func()
}

//...................................................................
var Pages []RepositoryMigrationPage

//========================================================================================================================
func AddMigration(repo RepositoryMigration) int {
	if len(Pages) == 0 || len(Pages[len(Pages)-1].page_repositories) == 1000 {
		var page RepositoryMigrationPage
		page.page_repositories = append(page.page_repositories, repo)
		Pages = append(Pages, page)
		return (len(Pages) - 1)
	}
	Pages[len(Pages)-1].page_repositories = append(Pages[len(Pages)-1].page_repositories, repo)
	return (len(Pages) - 1)
}

//========================================================================================================================

//========================================================================================================================
