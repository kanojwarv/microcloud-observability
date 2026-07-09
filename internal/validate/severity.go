package validate

// Severity indicates whether a validation check is mandatory.
type Severity int

const (
	Required Severity = iota
	Optional
)
