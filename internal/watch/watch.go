package watch

import (
	"log"
	"time"

	"github.com/kanojwarv/microcloud-observability/internal/generate"
)

func Start(
	repositoryRoot string,
) {

	ticker := time.NewTicker(
		5 * time.Second,
	)

	defer ticker.Stop()

	for range ticker.C {

		err := generate.All(
			repositoryRoot,
		)

		if err != nil {

			log.Println(
				err,
			)

			continue
		}

		log.Println(
			"Artifacts regenerated",
		)
	}
}
