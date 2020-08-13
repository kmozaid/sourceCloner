package model

type Config struct {
	Port         int16
	ClientId     string
	ClientSecret string
	RedirectUri  string
	Scopes       []string
	CloneDir     string
}

type User struct {
	UserName string `json:"login"`
	ReposUrl string `json:"repos_url"`
}

type Repository struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	CloneUrl string `json:"clone_url"`
}

type RepositoryList struct {
	AccessToken  string
	Repositories []Repository
}

type CloneResponse struct {
	Status string `json:"status"`
}
