package model

// Metric describes a metric definition.
type Metric struct {
	Metric   string   `yaml:"metric"`
	Object   Object   `yaml:"object"`
	Category Category `yaml:"category"`
	Unit     string   `yaml:"unit"`

	RecordingRule string   `yaml:"recording_rule"`
	Dashboards    []string `yaml:"dashboards"`
	Alerts        []string `yaml:"alerts"`
}

// MetricsFile represents a metrics YAML file.
type MetricsFile struct {
	Metrics []Metric `yaml:"metrics"`
}
