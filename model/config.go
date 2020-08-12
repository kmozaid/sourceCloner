package model

type Config struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	Scopes       []string
	CloneDir     string
}
