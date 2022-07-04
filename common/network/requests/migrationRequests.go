package common

import (
	"main/config"
	repository "main/core/repository"
)

type Service string

const (
	GIT    Service = "git"
	GITHUB Service = "github"
	GITEA  Service = "gitea"
	gitlab Service = "gitlab"
)

type MigrationRequest struct {
	Auth_password   string `json:"auth_password"`
	Auth_token      string `json:"auth_token"`
	Auth_username   string `json:"auth_username"`
	Clone_addr      string `json:"clone_addr"`
	Description     string `json:"description"`
	Issues          bool   `json:"issues"`
	Iabels          bool   `json:"labels"`
	Ifs             bool   `json:"lfs"`
	Ifs_endpoint    string `json:"lfs_endpoint"`
	Milestones      bool   `json:"milestones"`
	Mirror          bool   `json:"mirror"`
	Mirror_interval string `json:"mirror_interval"`
	Private         bool   `json:"private"`
	Pull_requests   bool   `json:"pull_requests"`
	Releases        bool   `json:"releases"`
	Repo_name       string `json:"repo_name"`
	Repo_owner      string `json:"repo_owner"`
	Service         string `json:"service"`
	Uid             int64  `json:"uid"`
	Wiki            bool   `json:"wiki"`
}
type MigrationReponse struct {
	Alow_merge_commits          bool   `json:"allow_merge_commits"`
	Allow_rebase                bool   `json:"allow_rebase"`
	Allow_rebase_explicit       bool   `json:"allow_rebase_explicit"`
	Allow_squash_merge          bool   `json:"allow_squash_merge"`
	Archived                    bool   `json:"archived"`
	Avatar_url                  string `json:"avatar_url"`
	Clone_url                   string `json:"clone_url"`
	Created_at                  string `json:"created_at"`
	Default_branch              string `json:"default_branch"`
	Default_merge_style         string `json:"default_merge_style"`
	Description                 string `json:"description"`
	Empty                       bool   `json:"empty"`
	Fork                        bool   `json:"fork"`
	Forks_count                 int64  `json:"forks_count"`
	Full_name                   string `json:"full_name"`
	Has_issues                  bool   `json:"has_issues"`
	Has_projects                bool   `json:"has_projects"`
	Has_pull_requests           bool   `json:"has_pull_requests"`
	Has_wiki                    bool   `json:"has_wiki"`
	Html_url                    string `json:"html_url"`
	Id                          int64  `json:"id"`
	Ignore_whitespace_conflicts bool   `json:"ignore_whitespace_conflicts"`
	Internal                    bool   `json:"internal"`
	Mirror                      bool   `json:"mirror"`
	Mirror_interval             string `json:"mirror_interval"`
	Name                        string `json:"name"`
	Open_issues_count           int64  `json:"open_issues_count"`
	Open_pr_counter             int64  `json:"open_pr_counter"`
	Original_url                string `json:"original_url"`
	Private                     bool   `json:"private"`
	Release_counter             int64  `json:"release_counter"`
	Size                        int64  `json:"size"`
	Ssh_url                     string `json:"ssh_url"`
	Stars_count                 int64  `json:"stars_count"`
	Template                    bool   `json:"template"`
	Updated_at                  string `json:"updated_at"`
	Watchers_count              int64  `json:"watchers_count"`
	Website                     string `json:"website"`
}

//..................................................................
func MakeMigrationRequest(migration repository.RepositoryMigration) Request {
	var baseRequest Request
	var headers []Header
	var header1 Header
	var header2 Header
	var header3 Header

	header1.Key = "Authorization"
	header1.Value = "token " + config.ConfigStruct.Config.AuthToken
	header2.Key = "Content-Type"
	header2.Value = "application/json"
	header3.Key = "cookies"
	header3.Value = "_csrf=Bcz_qZgF9DZXnJLTHSs-QGfsWho6MTY0NjMyMjk0NTc1ODUwNjg0Ng; i_like_gitea=b0a068734bae7fab"
	headers = append(headers, header1)
	headers = append(headers, header2)
	//headers = append(headers, header3)

	var body MigrationRequest
	body.Clone_addr = migration.Link
	body.Repo_name = migration.Name
	body.Repo_owner = config.ConfigStruct.Config.Owner
	body.Service = migration.Service

	baseRequest.Headers = headers
	baseRequest.Body = body
	baseRequest.Method = POST
	baseRequest.Url = config.ConfigStruct.Config.BaseAddress + "api/v1/repos/migrate"
	return baseRequest

}
