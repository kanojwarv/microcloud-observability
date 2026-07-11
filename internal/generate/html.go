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
    			font-family: Arial, sans-serif;
    			background-color: #f5f5f5;
    			padding: 40px;
			}

			h1 {
    			text-align: center;
    			margin-bottom: 40px;
			}

			.graph {
    			margin-bottom: 50px;
			}

			.node {
    			background: white;
    			border: 2px solid #333;
    			border-radius: 10px;
    			padding: 16px;
    			width: 320px;
    			margin: auto;
    			text-align: center;
    			font-size: 20px;
    			font-weight: bold;
			}

			.dependencies {
    			margin-top: 24px;
    			display: flex;
    			flex-direction: column;
    			align-items: center;
    			gap: 12px;
			}

			.dependency {
    			background: white;
    			border: 1px solid #888;
    			border-radius: 8px;
    			padding: 12px;
    			width: 280px;
    			text-align: center;
			}
    </style>
</head>

<body>

<h1>Dependency Graph</h1>
`

	for artifact, dependencies := range g.Dependencies() {

		html += fmt.Sprintf(`
        <div class="graph">

            <div class="node">
                %s
            </div>

            <div class="dependencies">
    `, artifact)

		for _, dependency := range dependencies {

			html += fmt.Sprintf(`
            <div class="dependency">
                %s
            </div>
        `, dependency)
		}

		html += `
            </div>

        </div>
    `
	}

	html += `
</body>
</html>
`

	return []byte(
		html,
	), nil
}
