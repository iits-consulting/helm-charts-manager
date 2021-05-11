package docopt

func getUsageInfo() string {
	return `Auto deployment service for helm-charts.

Usage:
  helm-charts-manager -h | --help
  helm-charts-manager -v | --version
  helm-charts-manager plan [options]
  helm-charts-manager apply [options]
  helm-charts-manager list-unmanaged [--skip-namespaces=namespace1,namespace2] [options]
  helm-charts-manager lint [options]
Options:
  -h --help                                 Show this screen.
  -v --version                              Show version.
  --config-file <configFile>                Configuration file for helm-charts-manager [default: ./helm-charts-manager-config.yaml]
  --charts-path <chartsPath>                Base path of the helm charts [default: ./]
  --charts <charts>                         List of charts to be planned/deployed, uses all charts in <configFile> if not specified.
  --auto-approve                            Automatically answer all queries with yes.
  --update                               	Update charts dependencies and Version.
  --debug                                   Enable debug printouts for every command executed.`
}
