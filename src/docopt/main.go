package docopt

import (
	"github.com/docopt/docopt-go"
	"strings"
)

// HelmChartsManagerArgs : command line argument template for the helm-charts-manager
type HelmChartsManagerArgs struct {
	Apply          bool
	Plan           bool
	Lint           bool
	ConfigFilePath string
	ChartsBasePath string
	ChartNames     []string
	AutoApprove    bool
	Update         bool
	Debug          bool
}

// ParseConfigFromArgs : parsing and handling of command line arguments
func ParseConfigFromArgs(version string) HelmChartsManagerArgs {
	arguments, _ := docopt.ParseArgs(getUsageInfo(), nil, version)
	var config HelmChartsManagerArgs
	config.Apply = arguments["apply"].(bool)
	config.Plan = arguments["plan"].(bool)
	config.Lint = arguments["lint"].(bool)

	config.ConfigFilePath = arguments["--config-file"].(string)
	config.ChartsBasePath = arguments["--charts-path"].(string)
	if arguments["--charts"] != nil {
		config.ChartNames = strings.Split(arguments["--charts"].(string), ",")
	}
	config.AutoApprove = arguments["--auto-approve"].(bool)
	config.Update = arguments["--update"].(bool)
	config.Debug = arguments["--debug"].(bool)
	return config
}
