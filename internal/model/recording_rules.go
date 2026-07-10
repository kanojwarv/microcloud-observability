package model

// RecordingRuleFile represents a Prometheus recording rule file.
type RecordingRuleFile struct {
	Groups []RuleGroup `yaml:"groups"`
}

// RuleGroup represents a Prometheus rule group.
type RuleGroup struct {
	Name     string          `yaml:"name"`
	Interval string          `yaml:"interval"`
	Rules    []RecordingRule `yaml:"rules"`
}

// RecordingRule represents a single recording rule.
type RecordingRule struct {
	Record string `yaml:"record"`
	Expr   string `yaml:"expr"`
}
