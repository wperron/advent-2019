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

func Edges(g Graph) map[Node][]Node {
	return g.edges
}

func AddEdge(g *Graph, from, to Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]Node)
	}
	if !contains(g.edges[from], to) {
		g.edges[from] = append(g.edges[from], to)
		g.edges[to] = append(g.edges[to], from)
	}
}

func Traverse(g Graph, start, end Node) int {
	distances := make(map[Node]int)
	distances[start] = 0
	queue := []Node{start}
	visited := []Node{queue[0]}

	for distances[end] == 0 {
		curr := queue[0]
		dist := distances[curr]
		rel := g.edges[curr]
		for _, adj := range rel {
			if contains(visited, adj) {
				continue
			}
			distances[adj] = dist + 1
			visited = append(visited, adj)
			queue = append(queue, adj)
		}
		queue = queue[1:]
	}

	return distances[end]
}

func Trace(g Graph, start, end Node) int {
	var hops int
	// magic maybe?
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
