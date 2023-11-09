package graphs

type WNode struct {
	vertex string
	weight int
}

type WeightedGraph struct {
	// Map each vertex to a list of weighted nodes it is connected to
	adjacencyList map[string][]WNode
}

// NewWeightedGraph creates a WeightedGraph from a given adjacency list
func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		adjacencyList: make(map[string][]WNode),
	}
}

// AddVertex adds a new vertex to the graph
func (g *WeightedGraph) AddVertex(vertex string) {
	if _, exists := g.adjacencyList[vertex]; !exists {
		g.adjacencyList[vertex] = []WNode{}
	}
}

// AddEdge adds a new edge with weight to the graph
// Since it's an undirected graph, we add an edge from src to dest and from dest to src
func (g *WeightedGraph) AddEdge(src, dest string, weight int) {
	g.adjacencyList[src] = append(g.adjacencyList[src], WNode{vertex: dest, weight: weight})
	g.adjacencyList[dest] = append(g.adjacencyList[dest], WNode{vertex: src, weight: weight})
}

// GetVertices returns a slice of all the graph's vertices
func (g *WeightedGraph) GetVertices() []string {
	vertices := make([]string, 0, len(g.adjacencyList))
	for vertex := range g.adjacencyList {
		vertices = append(vertices, vertex)
	}
	return vertices
}

// GetNeighbors returns a slice of all the weighted neighbors of a given vertex
func (g *WeightedGraph) GetNeighbors(vertex string) []WNode {
	return g.adjacencyList[vertex]
}
