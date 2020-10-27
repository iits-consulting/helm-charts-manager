package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func executeHelmLint(charts []string) {
	executeCommand(true, "helm", "lint --with-subcharts "+strings.Join(charts, " "))
}

func executeHelmCommandsForCharts(charts []string, helmCommand string) []string {
	var changedCharts []string
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(charts))
	for _, helmChartName := range charts {
		go executeHelmCommand(&waitGroup, &changedCharts, helmChartName, helmCommand)
	}
	waitGroup.Wait()
	return changedCharts
}

func executeHelmCommand(waitGroup *sync.WaitGroup, changedCharts *[]string, helmChartName string, helmCommand string) {
	helmCommand += determineAdditionalArgs(helmChartName)
	helmCommand = strings.ReplaceAll(helmCommand, "<chartName>", helmChartName)
	helmCommand = os.ExpandEnv(helmCommand)
	chartChanges := executeCommand(false, "helm", helmCommand)
	if len(chartChanges) > 0 {
		*changedCharts = append(*changedCharts, helmChartName)
		fmt.Println(chartChanges)
	}
	waitGroup.Done()
}

func determineAdditionalArgs(helmChartName string) string {
	overrideArgs := Config.Stages[Stage][helmChartName]
	defaultArgs := Config.DefaultHelmArgs[helmChartName]
	if len(overrideArgs) > 0 {
		return " " + overrideArgs
	} else if len(defaultArgs) > 0 {
		return " " + defaultArgs
	} else {
		return ""
	}
}
