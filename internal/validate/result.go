package validate

// Check represents a single validation result.
type Check struct {
	Name     string
	Passed   bool
	Severity Severity
	Message  string
}

// Result represents the outcome of one validator.
type Result struct {
	Name     string
	Checks   []Check
	Passed   bool
	Score    int
	MaxScore int
}
