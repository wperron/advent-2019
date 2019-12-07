package graph

type Node struct {
	Val string
}

type Graph struct {
	nodes []Node
	edges map[Node][]Node
}

func AddNode(g *Graph, n Node) {
	if !contains(g.nodes, n) {
		g.nodes = append(g.nodes, n)
	}
}

func Nodes(g Graph) []Node {
	return g.nodes
}

func AddEdge(g *Graph, from, to Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]Node)
	}
	if !contains(g.edges[from], to) {
		g.edges[from] = append(g.edges[from], to)
	}
}

func Traverse(g Graph, start Node) int {
	var paths int
	queue := []Node{start}

	for len(queue) > 0 {
		curr := queue[0]
		rel := g.edges[curr]
		for _, adj := range rel {
			paths++
			queue = append(queue, adj)
		}
		queue = queue[1:]
	}

	return paths
}

func Trace(g Graph, start, end Node) int {
	var hops int
	return hops
}

func contains(a []Node, x Node) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
