package generate

import (
	"encoding/json"

	"github.com/kanojwarv/microcloud-observability/internal/graph"
)

func GraphJSON(
	repositoryRoot string,
) ([]byte, error) {

	g, err := graph.Build(
		repositoryRoot,
	)
	if err != nil {
		return nil, err
	}

	data := map[string][]string{}

	for artifact, dependencies := range g.Dependencies() {

		data[artifact] = dependencies
	}

	return json.MarshalIndent(
		data,
		"",
		"  ",
	)
}
