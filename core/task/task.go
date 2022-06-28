package core

import (
	"main/config"
)

type JobStatus string

const (
	pending JobStatus = "pending"
	running JobStatus = "running"
	done    JobStatus = "done"
	failed  JobStatus = "failed"
)

type Task struct {
	taskId            int
	taskConfiguration config.Task
}

type Job struct {
	jobId     int
	jobConfig config.Config
	tasks     []Task
	status    JobStatus
}
