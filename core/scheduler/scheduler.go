package core

import (
	job "main/core/job"
)

//########################################################################################################################
//Function to schedule a job instance
func Schedule(job job.Job) {
	job.Run(job)
}

//========================================================================================================================
// Function to register the state of a job
func RegisterState() {}

//========================================================================================================================
// Function to control the activity of the logger
func ActivateLogs(log_level int) bool { return false }
