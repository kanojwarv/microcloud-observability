package graph

type Graph struct {
	dependencies map[string][]string
	dependents   map[string][]string
}

func New() *Graph {
	return &Graph{
		dependencies: make(map[string][]string),
		dependents:   make(map[string][]string),
	}
}

func (g *Graph) AddDependency(
	from string,
	to string,
) {

	g.dependencies[from] = append(
		g.dependencies[from],
		to,
	)

	g.dependents[to] = append(
		g.dependents[to],
		from,
	)
}

func (g *Graph) DependenciesOf(
	name string,
) []string {

	return g.dependencies[name]
}

func (g *Graph) DependentsOf(
	name string,
) []string {

	return g.dependents[name]
}
