package core

import (
	"main/config"
	core "main/core/repository"

	"github.com/google/uuid"
)

//########################################################################################################################
type TaskStatus string
type RunTask func(task Task) TaskStatus

const (
	pending TaskStatus = "pending"
	running TaskStatus = "running"
	done    TaskStatus = "done"
	failed  TaskStatus = "failed"
)

type Task struct {
	TaskId            uuid.UUID
	TaskConfiguration config.Task
	Page              core.RepositoryMigrationPage
	Run               RunTask
	Status            TaskStatus
}

//...................................................................
var Tasks []Task

//========================================================================================================================
/**
Function to create a new task according to a page of repositories and configuration file
*/
func NewTask(page core.RepositoryMigrationPage, config config.Task, runFunc RunTask) uuid.UUID {
	var newTask Task
	uuid := uuid.New()
	newTask.TaskId = uuid
	newTask.Page = page
	newTask.Run = runFunc
	newTask.TaskConfiguration = config
	Tasks = append(Tasks, newTask)
	return uuid
}
