package core

import (
	"fmt"
	logger "main/common/logger"
	network "main/common/network"
	networkRequests "main/common/network/requests"
	"main/config"
	job "main/core/job"
	parser "main/core/jsonParser"
	repository "main/core/repository"
	scheduler "main/core/scheduler"
	seeder "main/core/seeder"
	task "main/core/task"
	"net/http"
	"os"

	"github.com/jinzhu/configor"
)

//########################################################################################################################

func MigrationPostRequestCallBack(response *http.Response) error {

	var res networkRequests.MigrationReponse
	err := network.ParsResponse(response, &res)
	logger.LoggerReport.Level = logger.Info
	logger.LoggerReport.Message = res
	logger.PrintLog(logger.LoggerReport)

	return err

}
func MigrationCallBack(migration repository.RepositoryMigration) {

	request := networkRequests.MakeMigrationRequest(migration)

	request.Callback = MigrationPostRequestCallBack
	fmt.Print(request)
	err := network.Send(request)
	logger.LoggerReport.Detail = logger.MigrationReport{RepoName: migration.Name}
	logger.LoggerReport.Level = logger.Error
	if err != nil {
		logger.LoggerReport.Message = err.Error()
		logger.PrintLog(logger.LoggerReport)
	}

}

//========================================================================================================================
func TaskCallBack(task task.Task) task.TaskStatus {
	fmt.Println("task with UID", task.TaskId, "with ", len(task.Page.Page_repositories), " repos has been started\n-----------------------------------------")
	logger.LoggerReport.TaskId = task.TaskId
	for _, migration := range task.Page.Page_repositories {
		migration.Service = task.TaskConfiguration.Service
		migration.Run(migration)
	}

	return task.Status
}

//========================================================================================================================
func JobCallBack(job job.Job) job.JobStatus {
	logger.LoggerReport.JobId = job.JobId
	fmt.Println("job with UID", job.JobId, "with ", len(job.Tasks), " tasks has been started\n-----------------------------------------")
	for _, task := range job.Tasks {
		task.Run(task)
	}
	return job.Status
}

//========================================================================================================================
func InitConfig(configPath string) error {
	err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&config.ConfigStruct, configPath)
	return err
}

//========================================================================================================================
//handles start command with it's specified arguments
func StartHandler(configPath string, isService bool) (bool, error) {

	ConfigLogger()
	err := InitConfig(configPath)
	parser.ParseJson(config.ConfigStruct.Config.Tasks[0].Task.DataSet.File)
	job, _ := seeder.SeedAll(&config.ConfigStruct, parser.Dataset, MigrationCallBack, TaskCallBack, JobCallBack)

	scheduler.Schedule(job)

	return true, err

}

//========================================================================================================================
func ConfigLogger() {
	logger.LoggerReport.Level = logger.Info
	logger.LoggerReport.ProcessId = os.Getpid()
	logger.InitLogger(logger.LoggerOptions{PrintDestination: logger.File, FileAddress: "output.log", LogLevels: append(logger.LogOptions.LogLevels, logger.Info, logger.Error)})
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
