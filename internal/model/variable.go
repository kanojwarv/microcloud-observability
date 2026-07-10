package model

// Variable represents a dashboard variable.
type Variable struct {
	Name  string `yaml:"name"`
	Label string `yaml:"label"`
}

// VariablesFile represents the variables registry.
type VariablesFile struct {
	Variables []Variable `yaml:"variables"`
}
