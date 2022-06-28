package config

type ConfigTestFactory struct {
	base_url   string
	auth_token string
	owner      string
	Tasks      []TaskTestFactory
}
type TaskTestFactory struct {
	name    string
	service string
	DataSet DatasetTestFactory
}
type DatasetTestFactory struct {
	File     string
	MaxStars int
}

var ConfigFactory ConfigTestFactory

func InitFactory() {
	ConfigFactory.base_url = "https://https://git.techpal.ru/"
	ConfigFactory.auth_token = "8c151ab6d25380a8b3018a2f5c4a40ac8ec06836"
	ConfigFactory.owner = "comfortech"
	var task TaskTestFactory
	task.name = "github"
	task.service = "github"
	var dataset DatasetTestFactory
	dataset.File = "repos.json"
	dataset.MaxStars = 50
	task.DataSet = dataset
	ConfigFactory.Tasks = append(ConfigFactory.Tasks, task)

}
