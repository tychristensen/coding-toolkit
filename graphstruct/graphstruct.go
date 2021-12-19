package graphstruct

import "sync"

// Node is the data structure representing a single Node
type Node struct {
	Value string
}

// Graph is the data structure representing a graph of Nodes
type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

// String returns a string version of a node
func (n Node) String() string {
	return n.Value
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()

	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	g.lock.Unlock()
}

// String returns a string version of a graph
func (g *Graph) String() (str string) {
	str = "Nodes:\n"
	for _, node := range g.nodes {
		str += node.String() + "\n"
	}
	str += "\nEdges:\n"
	for n1, nodes := range g.edges {
		for _, n2 := range nodes {
			str += n1.String() + " -> " + n2.String() + "\n"
		}
		str += "\n"
	}

	return str
}
