package model

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// ReadConfigFile : reading, parsing and handling of the config file
func ReadConfigFile(configFilePath string) Config {
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	config := Config{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}
	return config
}
