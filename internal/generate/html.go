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

            cursor: pointer;

            transition: transform 0.2s ease;
        }

        .dependency:hover {
            transform: scale(
                1.05
            );
        }

        #impact-panel {
            position: fixed;

            top: 40px;
            right: 40px;

            width: 320px;

            background: white;

            border: 2px solid #333;

            border-radius: 10px;

            padding: 20px;

            box-shadow: 0 4px 12px rgba(
                0,
                0,
                0,
                0.15
            );
        }
		#doctor-panel {

    		position: fixed;

    		top: 320px;

    		right: 40px;

    		width: 320px;

    		background: white;

    		border: 2px solid #333;

    		border-radius: 10px;

    		padding: 20px;

    		box-shadow: 0 4px 12px rgba(
        		0,
        		0,
        		0,
        		0.15
    		);
		}
		
		#search-panel {

    	position: fixed;

    	top: 40px;

    	left: 40px;

    	width: 320px;

    	background: white;

    	border: 2px solid #333;

    	border-radius: 10px;

    	padding: 20px;
		}

    </style>

</head>

<body>

    <h1>Dependency Graph</h1>

    <div id="impact-panel">

        <h2>Impact Analysis</h2>

        <p>Select a metric...</p>

    </div>
		<div id="doctor-panel">

    	<h2>Repository Health</h2>

    	<p>Loading...</p>

	</div>
	<div id="search-panel">

    	<h2>Repository Search</h2>

    	<input
        	id="search-input"
        	placeholder="Search..."
    	>

    	<button
        	onclick="performSearch()"
    	>
        	Search
    	</button>
		<button
        id="resume-button"
        onclick="resumeRefresh()"
        style="display: none;"
    	>
        Resume Live Updates
        </button>

    	<div id="search-results">

        	<p>No search yet.</p>

    	</div>

</div>
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
    			<div
        		class="dependency"
        		onclick="showImpact('%s')"
    			>
        		%s
    			</div>
				`, dependency, dependency)
		}

		html += `
            </div>

        </div>
    `
	}

	html += `
		<script>

		let autoRefresh = true

		async function showImpact(
			artifact,
		) {

			const response = await fetch(
				"/api/impact/" + artifact,
			)

			if (!response.ok) {

				document.getElementById(
					"impact-panel",
				).innerHTML =
					"<h2>Impact Analysis</h2>" +
					"<p>Artifact not found.</p>"

				return
			}

			const data = await response.json()

			let html =
				"<h2>Impact Analysis</h2>" +
				"<p><strong>Artifact:</strong> " +
				data.artifact +
				"</p>" +
				"<h3>Referenced by</h3>" +
				"<ul>"

			for (
				const item
				of data.referencedBy
			) {

				html +=
					"<li>" +
					item +
					"</li>"
			}

			html += "</ul>"

			document.getElementById(
				"impact-panel",
			).innerHTML = html
		}

		async function refreshDoctor() {

			const response = await fetch(
				"/api/doctor",
			)

			const data = await response.json()

			let html =
				"<h2>Repository Health</h2>"

			for (
				const check
				of data.checks
			) {

				html +=
					"<p>" +
					check.status +
					" : " +
					check.name +
					"</p>"

				html +=
					"<small>" +
					check.message +
					"</small>"
			}

			document.getElementById(
				"doctor-panel",
			).innerHTML = html
		}
		
		refreshDoctor()

		async function performSearch() {

			autoRefresh = false

			document.getElementById(
    		"resume-button",
			).style.display = "block"

    		const query = document.getElementById(
        	"search-input",
    		).value

    		const response = await fetch(
        	"/api/search?q=" + query,
    		)

    		const data = await response.json()

    		let html = ""

    		for (
        		const result
        		of data.results
    		) {

        	html +=
            	"<p><strong>" +
            	result.kind +
            	"</strong></p>"

        	html +=
            	"<p>" +
            	result.name +
            	"</p>"

        	html +=
            	"<small>" +
            	result.path +
            	"</small><hr>"
    		}

    		document.getElementById(
        	"search-results",
    		).innerHTML = html
		}
		
		function resumeRefresh() {

    		autoRefresh = true

    		document.getElementById(
        	"resume-button",
    		).style.display = "none"

   			 window.location.reload()
		}


		setInterval(
    		function () {

        	if (autoRefresh) {

            	window.location.reload()
        	}

    	},
    	5000,
		)

</script>

</body>

</html>
`

	return []byte(
		html,
	), nil
}
