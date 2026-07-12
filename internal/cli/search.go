package cli

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/search"
)

func executeSearch(
	query string,
) error {

	response, err := search.Find(
		".",
		query,
	)

	if err != nil {
		return err
	}

	fmt.Println(
		"Search Results",
	)

	fmt.Println(
		"==============",
	)

	fmt.Println()

	for _, result := range response.Results {

		fmt.Printf(
			"[%s]\n\n",
			result.Kind,
		)

		fmt.Printf(
			"    %s\n",
			result.Name,
		)

		fmt.Printf(
			"    %s\n\n",
			result.Path,
		)
	}

	return nil

}
