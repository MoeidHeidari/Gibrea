package core

import (
	"main/config"
	job "main/core/job"
	parser "main/core/jsonParser"
	repository "main/core/repository"
	task "main/core/task"
)

//########################################################################################################################
func SeedMigrations(repos []parser.Repos, functionality repository.MigrationRunFunc) {
	for i := 0; i < len(repos); i++ {
		migration := repository.NewMigration(repos[i].Name, repos[i].Link, repos[i].Description, repos[i].Stars, functionality)
		repository.AddMigration(migration)
	}
}

//========================================================================================================================
func SeedTasks(pages []repository.RepositoryMigrationPage, config config.Task, functionality task.RunTask) {
	for i := 0; i < len(pages); i++ {
		task.NewTask(pages[i], config, functionality)
	}
}

//========================================================================================================================
func SeedJob(tasks []task.Task, functionality job.RunJob) job.Job {
	job := job.NewJob(tasks, functionality)
	return job
}

//========================================================================================================================

func SeedAll(config *config.ConfigStructure, data parser.RepoList, migration_functionality repository.MigrationRunFunc, task_functionality task.RunTask, job_functionality job.RunJob) (job.Job, error) {
	for _, page := range data.Pages {
		SeedMigrations(page.Repos, migration_functionality)
	}

	SeedTasks(repository.Pages, config.Config.Tasks[0].Task, task_functionality)
	job := SeedJob(task.Tasks, job_functionality)
	return job, nil
}
