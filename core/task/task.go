package core

import (
	"main/config"
	core "main/core/repository"

	"github.com/google/uuid"
)

//########################################################################################################################
type Task struct {
	taskId            uuid.UUID
	taskConfiguration config.Task
	page              core.RepositoryMigrationPage
}

//...................................................................
var Tasks []Task

//========================================================================================================================
/**
Function to create a new task according to a page of repositories and configuration file
*/
func NewTask(page core.RepositoryMigrationPage, config config.Task) uuid.UUID {
	var newTask Task
	uuid := uuid.New()
	newTask.taskId = uuid
	newTask.page = page
	newTask.taskConfiguration = config
	Tasks = append(Tasks, newTask)
	return uuid
}
