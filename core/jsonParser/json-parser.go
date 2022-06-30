package core

import (
	"encoding/json"
	"io/ioutil"
)

//########################################################################################################################
type RepoList struct {
	Pages []Pages `json:"pages"`
}
type Repos struct {
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
}
type Pages struct {
	Repos []Repos `json:"repos"`
}

//...................................................................
var Dataset RepoList

//========================================================================================================================
func ParseJson(path string) error {
	repos, _ := ioutil.ReadFile(path)
	err := json.Unmarshal(repos, &Dataset)
	return err
}
