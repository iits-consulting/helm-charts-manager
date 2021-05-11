package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func updateCharts(charts []string) {
	executeCommand(false, "helm", "repo update")
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(charts))
	for _, ChartName := range charts {
		go packageAndUpdateChart(&waitGroup, ChartName)
	}
	waitGroup.Wait()
}

//update chart version based on git version and tag and update dependencies
func packageAndUpdateChart(waitGroup *sync.WaitGroup, chartName string) {
	var outStrings []string
	executeCommand(false, "helm", "dependency update --skip-refresh "+chartName)
	outStrings = append(outStrings, executeCommand(false, "helm",
		"package "+chartName+" --version "+ReleaseVersion))
	executeCommand(false, "helm",
		"chart save "+chartName+"-"+ReleaseVersion+".tgz "+chartName)
	executeCommand(false, "helm",
		"chart export "+chartName+":"+ReleaseVersion)
	err := os.Remove(chartName + "-" + ReleaseVersion + ".tgz")
	if err != nil {
		outStrings = append(outStrings, err.Error())
	}
	fmt.Println(strings.Join(outStrings, "\n"))
	waitGroup.Done()
}
