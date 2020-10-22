package model

// Config : configuration file model
type Config struct {
	Environment     []Variable      `yaml:"env"`
	CommandSpec     CommandSpec     `yaml:"commandSpec"`
	DefaultHelmArgs DefaultHelmArgs `yaml:"defaultHelmArgs"`
	Stages          Stages          `yaml:"stages"`
}

// EnvironmentVariables : config file template for environment variables
type Variable map[string]string

// CommandSpec : config file template for command arguments
type CommandSpec map[string]string

// DefaultHelmArgs : config file template for default helm arguments
type DefaultHelmArgs map[string]string

// Stages : config file template for deployment stages in config file
type Stages map[string]Chart

// Chart : config file template for helm charts and argument overrides
type Chart map[string]string
