package doctor

type Check struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Report struct {
	Status string  `json:"status"`
	Checks []Check `json:"checks"`
}
