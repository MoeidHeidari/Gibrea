package core

import (
	"main/config"

	job "main/core/job"
	parser "main/core/jsonParser"
	repository "main/core/repository"
	task "main/core/task"
)

func Schedule(config config.Config, repos parser.RepoList) {

}

func SeedMigrations(repos []parser.Repos, functionality repository.MigrationRunFunc) {
	for i := 0; i < len(repos); i++ {
		migration := repository.NewMigration(repos[i].Name, repos[i].Link, repos[i].Description, repos[i].Stars, functionality)
		repository.AddMigration(migration)
	}
}
func InitTasks(pages []repository.RepositoryMigrationPage, config config.Task, functionality task.RunTask) {
	for i := 0; i < len(pages); i++ {
		task.NewTask(pages[i], config, functionality)
	}
}
func InitJob(tasks []task.Task, functionality job.RunJob) {
	job.NewJob(tasks, functionality)
}
func RegisterState()                  {}
func ActivateLogs(log_level int) bool { return false }
