package graph

import "testing"

func TestGraph(t *testing.T) {

	graph := New()

	graph.AddDependency(
		"executive",
		"vm_cpu_usage_seconds_rate",
	)

	dependencies := graph.DependenciesOf(
		"executive",
	)

	if len(dependencies) != 1 {
		t.Fatalf(
			"expected 1 dependency, got %d",
			len(dependencies),
		)
	}

	if dependencies[0] !=
		"vm_cpu_usage_seconds_rate" {

		t.Fatal(
			"unexpected dependency",
		)
	}
}
