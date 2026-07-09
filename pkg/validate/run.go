package validate

import "fmt"

func Run(root string) error {
	r := RepositoryValidator{Root: root}
	res := r.Validate()

	fmt.Println("MicroCloud Observability")
	fmt.Println("========================")
	fmt.Println("Repository Validation")

	for _, c := range res.Checks {
		mark := "✘"
		if c.Passed {
			mark = "✔"
		}

		fmt.Printf("%s %s\n", mark, c.Name)
	}

	if !res.Passed {
		return fmt.Errorf("repository validation failed")
	}

	fmt.Println()
	fmt.Println("SUCCESS")

	return nil
}
