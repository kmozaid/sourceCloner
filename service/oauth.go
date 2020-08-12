package service

import (
	"github.com/mozaidk/sourceCloner/service/provider"
)

func AuthorizeURL() string {
	return provider.AuthorizeURL()
}

func AccessToken(code string) string {
	return provider.AccessToken(code)
}
