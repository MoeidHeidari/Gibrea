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

var pages []RepositoryMigrationPage

func AddMigration(repo RepositoryMigration) int {
	if len(pages) == 0 || len(pages[len(pages)-1].page_repositories) == 1000 {
		var page RepositoryMigrationPage
		page.page_repositories = append(page.page_repositories, repo)
		pages = append(pages, page)
		return (len(pages) - 1)
	}
	pages[len(pages)-1].page_repositories = append(pages[len(pages)-1].page_repositories, repo)
	return (len(pages) - 1)
}
