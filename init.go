package main

import (
	"fmt"
	"helm-charts-manager/docopt"
	"helm-charts-manager/model"
	"log"
	"os"
	"strings"
)

// HelmChartsManagerArgs : command line arguments for the helm-charts-manager
var HelmChartsManagerArgs docopt.HelmChartsManagerArgs

// Config : helm-charts-manager parameters obtained from the config file
var Config model.Config

// Stage : stage obtained from the kubectl context
var Stage string

// ReleaseVersion : version number for the current deployment
var ReleaseVersion string

// PLAN : helm arguments for planning the chart
var PLAN string

// APPLY : helm arguments for applying the chart
var APPLY string

func init() {
	HelmChartsManagerArgs = docopt.ParseConfigFromArgs(VERSION)
	Config = model.ReadConfigFile(HelmChartsManagerArgs.ConfigFilePath)
	err := os.Chdir(HelmChartsManagerArgs.ChartsBasePath)
	if err != nil {
		panic(err)
	}
	PLAN = Config.CommandSpec["planCommandArgs"]
	APPLY = Config.CommandSpec["applyCommandArgs"]
	fmt.Println("Initializing environment...")
	initializeEnvironmentVariables()
	updateValues()
	fmt.Println("Environment initialization completed!")
}

func initializeEnvironmentVariables() {
	for _, variable := range Config.Environment {
		var value string
		for varName, varCommand := range variable {
			if _, exists := os.LookupEnv(varName); exists {
				fmt.Println("Skipping existing environment variable:", varName)
				continue
			}
			for _, command := range strings.Split(varCommand, " && ") {
				command = os.ExpandEnv(command)
				commandString := strings.SplitN(command, " ", 2)
				value = executeCommand(false, commandString[0], commandString[1])
			}
			err := os.Setenv(varName, value)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func updateValues() {
	ReleaseVersion = getEnvVariable("RELEASE_VERSION")
	Stage = getEnvVariable("STAGE")
}

func getEnvVariable(key string) string {
	envVariable := os.Getenv(key)
	if envVariable == "" {
		panic(key + " Environment variable is mandatory !")
	}
	return envVariable
}
