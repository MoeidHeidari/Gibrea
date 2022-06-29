package core

import (
	"main/config"

	job "main/core/job"
	parser "main/core/parser"
)

func Schedule(config config.Config, repos parser.RepoList) job.Job {
	InitMigrations()

}

func InitMigrations() {

}
func InitTasks() {}
func InitJob()   {}
func RegisterState()
func ActivateLogs(log_level int) bool { return false }
