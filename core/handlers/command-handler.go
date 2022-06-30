package core

import (
	"fmt"
	"main/config"
	job "main/core/job"
	parser "main/core/jsonParser"
	repository "main/core/repository"
	scheduler "main/core/scheduler"
	seeder "main/core/seeder"
	task "main/core/task"

	"github.com/jinzhu/configor"
)

//########################################################################################################################
//handles start command with it's specified arguments
func StartHandler(configPath string, isService bool) (bool, error) {

	err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&config.ConfigStruct, configPath)

	if err != nil {
		return false, err
	}
	parser.ParseJson(config.ConfigStruct.Config.Tasks[0].Task.DataSet.File)
	job, _ := seeder.SeedAll(&config.ConfigStruct, parser.Dataset, func(migration repository.RepositoryMigration) {

		fmt.Println("migration functionality")
	}, func(task task.Task) task.TaskStatus {
		fmt.Println("task with UID", task.TaskId, "with ", len(task.Page.Page_repositories), " repos has been started\n-----------------------------------------")
		for _, migration := range task.Page.Page_repositories {
			migration.Run(migration)
		}
		return task.Status
	}, func(job job.Job) job.JobStatus {

		fmt.Println("job with UID", job.JobId, "with ", len(job.Tasks), " tasks has been started\n-----------------------------------------")
		for _, task := range job.Tasks {
			task.Run(task)
		}
		return job.Status

	})
	scheduler.Schedule(job)

	return true, err

}

//========================================================================================================================
//handles stop command with it's specified arguments
func StopHandler() {
	fmt.Println("stop handler")
}

//========================================================================================================================
//handles pause command with it's specified arguments
func PauseHandler() {
	fmt.Println("pause handler")
}

//========================================================================================================================
//handles resum command with it's specified arguments
func ResumHandler() {
	fmt.Println("resum handler")
}

//========================================================================================================================
//handles status command with it's specified arguments
func StatusHandler() {
	fmt.Println("status handler")
}

//========================================================================================================================
//handles logs command with it's specified arguments
func LogsHandler(logFilePath string) {
	fmt.Println("logs handler", logFilePath)
}
