package core

import (
	"errors"
	"main/config"
	core "main/core/task"
	"os"

	"github.com/google/uuid"
)

//########################################################################################################################
type JobStatus string

var ErrNotFound = errors.New("Resource was not found")

type RunJob func(job Job) JobStatus

//...................................................................
//Job statuses
const (
	Pending JobStatus = "pending"
	Running JobStatus = "running"
	Done    JobStatus = "done"
	Failed  JobStatus = "failed"
)

//...................................................................
// Structure of a Job
type Job struct {
	JobId     uuid.UUID
	ProcessId int
	JobConfig config.Config
	Tasks     []core.Task
	Status    JobStatus
	Run       RunJob
}

var Jobs []Job

//========================================================================================================================
// Function to make a new job instance
func NewJob(tasks []core.Task, runFunc RunJob) Job {
	var newJob Job
	uid := uuid.New()
	newJob.JobId = uid
	newJob.Run = runFunc
	newJob.Status = Running
	newJob.Tasks = tasks
	newJob.JobConfig = config.ConfigStruct.Config
	newJob.ProcessId = os.Getpid()
	Jobs = append(Jobs, newJob)
	return newJob
}

//========================================================================================================================
// Function to return back a job associated with a UID
func getJob(uid uuid.UUID) (*Job, error) {
	for _, job := range Jobs {
		if job.JobId == uid {
			return &job, nil
		}
	}
	return nil, ErrNotFound
}
