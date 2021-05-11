package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type HelmList struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

func executeListUnmanaged(newCharts []string, namespacesToBeSkipped []string) {
	currentDeployedCharts := executeCommand(false, "helm", "list -A -o yaml")

	chartsToBeInspected := determineChartsToBeInspected(currentDeployedCharts, namespacesToBeSkipped)

	fmt.Println("Warning this charts are not managed by helm-charts-manager:")
	for _, oldChart := range chartsToBeInspected {
		if !containsString(newCharts, oldChart.Name) {
			fmt.Println("	Namespace:" + oldChart.Namespace + ", Name:" + oldChart.Name)
		}
	}
}

func determineChartsToBeInspected(currentDeployedCharts string, namespacesToBeSkipped []string) []HelmList {
	var chartsToBeInspected []HelmList
	var deployedCharts []HelmList
	_ = yaml.Unmarshal([]byte(currentDeployedCharts), &deployedCharts)

	for _, chart := range deployedCharts {
		if !containsString(namespacesToBeSkipped, chart.Namespace) {
			chartsToBeInspected = append(chartsToBeInspected, chart)
		}
	}
	return chartsToBeInspected
}

func containsString(stringArray []string, itemToSearchFor string) bool {
	var contains = false
	for _, oldChart := range stringArray {
		if oldChart == itemToSearchFor {
			contains = true
			break
		}
	}
	return contains
}
