package config

type ConfigStructure struct {
	Config Config `yaml:"config"`
}
type DataSet struct {
	File     string `yaml:"file"`
	MaxStars int    `yaml:"max_stars"`
}
type MultiProcessing struct {
	Enabled           bool `yaml:"enabled"`
	NumberOfProcesses int  `yaml:"number_of_processes"`
}
type ContinuesUpdate struct {
	Enabled         bool            `yaml:"enabled"`
	Interval        string          `yaml:"interval"`
	MultiProcessing MultiProcessing `yaml:"multi_processing"`
}
type SecuredData struct {
	Enabled    interface{} `yaml:"enabled"`
	AuthTokens []string    `yaml:"auth_tokens"`
	Issue      bool        `yaml:"issue"`
	Wikis      bool        `yaml:"wikis"`
	Branche    bool        `yaml:"branche"`
	Release    bool        `yaml:"release"`
	Label      bool        `yaml:"label"`
	Milstones  bool        `yaml:"milstones"`
}
type Task struct {
	Name            string          `yaml:"name"`
	Service         string          `yaml:"service"`
	DataSet         DataSet         `yaml:"data_set"`
	ContinuesUpdate ContinuesUpdate `yaml:"continues_update"`
	AsPrivate       bool            `yaml:"as_private"`
	SecuredData     SecuredData     `yaml:"secured_data"`
}
type Tasks struct {
	Task Task `yaml:"task"`
}
type Config struct {
	BaseAddress string  `yaml:"base_address"`
	AuthToken   string  `yaml:"auth_token"`
	Owner       string  `yaml:"owner"`
	Tasks       []Tasks `yaml:"tasks"`
}

var ConfigStruct ConfigStructure
