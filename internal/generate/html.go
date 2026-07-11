package generate

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/graph"
)

func GraphHTML(
	repositoryRoot string,
) ([]byte, error) {

	g, err := graph.Build(
		repositoryRoot,
	)

	if err != nil {
		return nil, err
	}

	html := `
<!DOCTYPE html>
<html>
<head>
    <title>MicroCloud Observability</title>

    <style>
        body {
            font-family: sans-serif;
            margin: 40px;
        }

        h1 {
            margin-bottom: 30px;
        }

        ul {
            margin-bottom: 20px;
        }
    </style>
</head>

<body>

<h1>Dependency Graph</h1>
`

	for artifact, dependencies := range g.Dependencies() {

		html += fmt.Sprintf(
			"<h2>%s</h2><ul>",
			artifact,
		)

		for _, dependency := range dependencies {

			html += fmt.Sprintf(
				"<li>%s</li>",
				dependency,
			)
		}

		html += "</ul>"
	}

	html += `
</body>
</html>
`

	return []byte(
		html,
	), nil
}
