package graphs

// UGraph is a representation of a simple unweighted graph using an adjacency list
type UGraph struct {
	// Map each vertex to a list of vertices it is connected to
	adjacencyList map[string][]string
}

// NewUnweightedGraph creates a UGraph from a given adjacency list
func NewUnweightedGraph() *UGraph {
	return &UGraph{
		adjacencyList: make(map[string][]string),
	}
}

// AddVertex adds a new vertex to the graph
func (g *UGraph) AddVertex(vertex string) {
	if _, exists := g.adjacencyList[vertex]; !exists {
		g.adjacencyList[vertex] = []string{}
	}
}

// sliceContains helper function to check if a slice contains a string
func sliceContains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// AddEdge adds a new edge to the graph
// Since it's an undirected graph, we add an edge from src to dest and from dest to src
func (g *UGraph) AddEdge(src, dest string) {
	// Ensure both the source and destination vertices exist
	if _, exists := g.adjacencyList[src]; !exists {
		g.AddVertex(src)
	}
	if _, exists := g.adjacencyList[dest]; !exists {
		g.AddVertex(dest)
	}

	// Add the edge from src to dest, if it doesn't already exist
	if !sliceContains(g.adjacencyList[src], dest) {
		g.adjacencyList[src] = append(g.adjacencyList[src], dest)
	}

	// Since it's an undirected graph, also add the edge from dest to src, if it doesn't already exist
	if !sliceContains(g.adjacencyList[dest], src) {
		g.adjacencyList[dest] = append(g.adjacencyList[dest], src)
	}
}

// GetVertices returns a slice of all the graph's vertices
func (g *UGraph) GetVertices() []string {
	vertices := make([]string, 0, len(g.adjacencyList))
	for vertex := range g.adjacencyList {
		vertices = append(vertices, vertex)
	}
	return vertices
}

// GetNeighbors returns a slice of all the neighbors of a given vertex
func (g *UGraph) GetNeighbors(vertex string) []string {
	return g.adjacencyList[vertex]
}

// DepthFirstSearch implements a simple Depth-First Search (DFS) into the graph and returns all the possible combinations
func (g *UGraph) DepthFirstSearch(startVertex string) []string {
	visited := make(map[string]bool)
	var output []string

	var dfs func(vertex string)
	dfs = func(vertex string) {
		visited[vertex] = true
		output = append(output, vertex)

		for _, neighbor := range g.adjacencyList[vertex] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(startVertex)
	return output
}

// BreadthFirstSearch implements a simple Breadth-First Search (BFS) into the graph and returns all the possible combinations
func (g *UGraph) BreadthFirstSearch(startVertex string) []string {
	visited := make(map[string]bool)
	queue := []string{startVertex} // Use a slice as a queue
	var output []string

	visited[startVertex] = true

	for len(queue) > 0 {
		// Dequeue an element from the queue
		vertex := queue[0]
		queue = queue[1:]

		// Process the current node
		output = append(output, vertex)

		// Enqueue all unvisited neighbors
		for _, neighbor := range g.adjacencyList[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return output
}
