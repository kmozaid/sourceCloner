package model

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
	Repositories []Repository
}

type CloneResponse struct {
	Status string `json:"status"`
}
