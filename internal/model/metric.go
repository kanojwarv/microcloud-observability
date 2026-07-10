package model

type Metric struct {
	Name        string
	Description string
	Unit        string

	Category Category
	Object   Object

	Labels []string
}
