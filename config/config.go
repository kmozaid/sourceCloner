package config

import (
	"encoding/json"
	"github.com/mozaidk/sourceCloner/model"
	"io/ioutil"
	"log"
)

var ServiceConf model.Config = loadConfiguration("config.json")

func loadConfiguration(file string) model.Config {
	var serviceConf model.Config
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &serviceConf)
	if err != nil {
		log.Fatal(err)
	}
	return serviceConf
}
