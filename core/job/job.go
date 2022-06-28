package core

import (
	"main/config"
	core "main/core/task"
	"os"

	"github.com/google/uuid"
)

//########################################################################################################################
type JobStatus string

type RunJob func(job Job) JobStatus

//...................................................................
const (
	pending JobStatus = "pending"
	running JobStatus = "running"
	done    JobStatus = "done"
	failed  JobStatus = "failed"
)

//...................................................................
type Job struct {
	jobId     uuid.UUID
	processId int
	jobConfig config.Config
	tasks     []core.Task
	status    JobStatus
	Run       RunJob
}

var Jobs []Job

//========================================================================================================================
func NewJob(tasks []core.Task, runFunc RunJob) uuid.UUID {
	var newJob Job
	uid := uuid.New()
	newJob.jobId = uid
	newJob.Run = runFunc
	newJob.status = running
	newJob.tasks = tasks
	newJob.jobConfig = config.ConfigStruct.Config
	newJob.processId = os.Getpid()
	Jobs = append(Jobs, newJob)
	return uid
}

//========================================================================================================================
