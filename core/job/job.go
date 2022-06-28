package core

import (
	"main/config"
	core "main/core/task"
	"os"

	"github.com/google/uuid"
)

type JobStatus string

const (
	pending JobStatus = "pending"
	running JobStatus = "running"
	done    JobStatus = "done"
	failed  JobStatus = "failed"
)

type Job struct {
	jobId     uuid.UUID
	processId int
	jobConfig config.Config
	tasks     []core.Task
	status    JobStatus
}

func newJob(tasks []core.Task) uuid.UUID {
	var newJob Job
	uid := uuid.New()
	newJob.jobId = uid
	newJob.status = running
	newJob.tasks = tasks
	newJob.jobConfig = config.ConfigStruct.Config
	newJob.processId = os.Getpid()
	return uid
}
