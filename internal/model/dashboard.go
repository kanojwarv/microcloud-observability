package model

// DashboardFile represents a dashboard definition.
type DashboardFile struct {
	Dashboard Dashboard `yaml:"dashboard"`
}

// Dashboard represents a dashboard.
type Dashboard struct {
	Name      string    `yaml:"name"`
	Title     string    `yaml:"title"`
	Variables []string  `yaml:"variables"`
	Sections  []Section `yaml:"sections"`
	Source    string    `yaml:"-"`
}

// Section represents a logical grouping of panels.
type Section struct {
	Title  string  `yaml:"title"`
	Panels []Panel `yaml:"panels"`
}

// Panel represents a visualization.
type Panel struct {
	Title  string `yaml:"title"`
	Metric string `yaml:"metric"`
	Type   string `yaml:"type"`
}
