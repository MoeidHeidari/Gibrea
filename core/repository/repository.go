package core

//########################################################################################################################
//Migration Functionality signature
type MigrationRunFunc func(migration RepositoryMigration)

//...................................................................
//Structure of a repository Migration page
type RepositoryMigrationPage struct {
	Page_repositories []RepositoryMigration
}

//...................................................................
// Structure of a Repository migration
type RepositoryMigration struct {
	Name        string
	Link        string
	Description string
	Stars       int
	Run         MigrationRunFunc
}

//...................................................................
var Pages []RepositoryMigrationPage

//========================================================================================================================
// Function to add a migration to a page
func AddMigration(repo RepositoryMigration) int {
	if len(Pages) == 0 || len(Pages[len(Pages)-1].Page_repositories) == 1000 {
		var page RepositoryMigrationPage
		page.Page_repositories = append(page.Page_repositories, repo)
		Pages = append(Pages, page)
		return (len(Pages) - 1)
	}
	Pages[len(Pages)-1].Page_repositories = append(Pages[len(Pages)-1].Page_repositories, repo)
	return (len(Pages) - 1)
}

//========================================================================================================================
// Function to instantiate a new Migration
func NewMigration(name string, link string, description string, stars int, runFunc MigrationRunFunc) RepositoryMigration {
	var repo RepositoryMigration
	repo.Name = name
	repo.Link = link
	repo.Description = description
	repo.Stars = stars
	repo.Run = runFunc
	return repo
}
