package main

import (
	"fmt"
	"os"
	"strings"
)

// VERSION : helm-charts-manager version, automatically updated from the latest git tag during build time.
var VERSION = "undefined"

func main() {
	var helmCharts = createHelmChartsList()
	if HelmChartsManagerArgs.Update {
		fmt.Println("Updating helm charts...")
		updateCharts(helmCharts)
		fmt.Println("Updating helm charts completed!")
	}
	if HelmChartsManagerArgs.Lint {
		executeHelmLint(helmCharts)
	}
	if HelmChartsManagerArgs.Plan {
		fmt.Println("Planning helm charts...")
		chartsToChange := executeHelmCommandsForCharts(helmCharts, PLAN)
		fmt.Println("Planning completed!")
		fmt.Println("These Charts will be updated: " + strings.Join(chartsToChange, ","))
	}
	if HelmChartsManagerArgs.Apply {
		fmt.Println("Planning helm charts...")
		chartsToChange := executeHelmCommandsForCharts(helmCharts, PLAN)
		fmt.Println("Planning completed!")
		fmt.Println("These Charts will be updated: " + strings.Join(chartsToChange, ","))
		if shouldProcessBeContinued("Would you like to deploy these changes?") {
			fmt.Println("Deploying helm charts...")
			executeHelmCommandsForCharts(chartsToChange, APPLY)
			fmt.Println("Deployment completed!")
		}
	}
}

func createHelmChartsList() []string {
	var helmChartsMap = make(map[string]bool)
	var helmCharts []string
	if len(HelmChartsManagerArgs.ChartNames) > 0 {
		helmCharts = HelmChartsManagerArgs.ChartNames
	} else if !HelmChartsManagerArgs.Lint {
		for chart := range Config.Stages[Stage] {
			helmCharts = append(helmCharts, chart)
		}
	} else {
		for stage := range Config.Stages {
			for chart := range Config.Stages[stage] {
				helmChartsMap[chart] = true
			}
		}
		for chart := range helmChartsMap {
			helmCharts = append(helmCharts, chart)
		}
	}
	return helmCharts
}

func shouldProcessBeContinued(queryString string) bool {
	if HelmChartsManagerArgs.AutoApprove {
		return true
	}
	var answer string
	var validAnswers = []string{"y", "yes", "n", "no", "q", "quit"}
	fmt.Printf(queryString + " [Yes/No/Quit] : ")
	_, err := fmt.Scanf("%s", &answer)
	if err != nil {
		panic(err)
	}
	answer = strings.ToLower(answer)
	for index, a := range validAnswers {
		if answer == a {
			if index < 2 {
				return true
			} else if index < 4 {
				return false
			} else {
				os.Exit(0)
			}
		}
	}
	panic(answer + "is not a valid answer!")
}
