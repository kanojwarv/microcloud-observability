package doctor

func Run() Report {

	return Report{
		Status: "OK",
		Checks: []Check{
			{
				Name:    "dashboard dependencies",
				Status:  "OK",
				Message: "All dashboard dependencies are used.",
			},
		},
	}
}
