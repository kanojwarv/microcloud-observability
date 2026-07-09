package validate

import "os"

type RepositoryValidator struct{ Root string }

func (r RepositoryValidator) Validate() Result {
	req := []string{"docs", "metrics", "recording-rules", "variables", "dashboards", "alerts"}
	res := Result{Name: "Repository"}
	okAll := true
	for _, p := range req {
		_, err := os.Stat(r.Root + "/" + p)
		ok := err == nil
		if !ok {
			okAll = false
		}
		res.Checks = append(res.Checks, Check{Name: p, Passed: ok})
	}
	res.Passed = okAll
	return res
}
