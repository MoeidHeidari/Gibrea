package core

import (
	"main/config"
	core "main/core/repository"

	"github.com/google/uuid"
)

type Task struct {
	taskId            uuid.UUID
	taskConfiguration config.Task
	page              core.RepositoryMigrationPage
}

type TaskLog struct {
	taskId  uuid.UUID
	jobId   uuid.UUID
	message string
}

func newTask(page core.RepositoryMigrationPage, config config.Task) uuid.UUID {
	var newTask Task
	uuid := uuid.New()
	newTask.taskId = uuid
	newTask.page = page
	newTask.taskConfiguration = config
	tasks = append(tasks, newTask)
	return uuid
}

var tasks []Task
