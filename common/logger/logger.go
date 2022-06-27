package common

import "github.com/xtgo/uuid"

type JobLog struct {
	jobId     uuid.UUID
	message   int
	processId int
}
