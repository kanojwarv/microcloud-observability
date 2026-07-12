package search

type Result struct {
	Kind string `json:"kind"`

	Name string `json:"name"`

	Path string `json:"path"`
}

type Response struct {
	Query string `json:"query"`

	Results []Result `json:"results"`
}
