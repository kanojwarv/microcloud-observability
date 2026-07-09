package validate

type Check struct {
	Name   string
	Passed bool
}
type Result struct {
	Name   string
	Passed bool
	Checks []Check
}
