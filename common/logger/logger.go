package common

import "github.com/xtgo/uuid"

type JobReport struct {
	JobId     uuid.UUID
	Message   int
	ProcessId int
}

//...................................................................
type TaskReprot struct {
	taskId  uuid.UUID
	jobId   uuid.UUID
	message string
}
