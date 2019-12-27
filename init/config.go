package init

import (
	"io/ioutil"
	"log"
	"oauth_lambda/model"

	"gopkg.in/yaml.v2"
)

// Conf is init config file
func Conf() model.Config {
	c := model.Config{}
	yamlFile, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return c
}
