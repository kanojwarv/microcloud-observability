package validate

// Report aggregates all validator results.
type Report struct {
	Results  []Result
	Passed   bool
	Score    int
	MaxScore int
}
