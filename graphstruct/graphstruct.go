package graphstruct

import "sync"

// Node is the data structure representing a single Node
type Node struct {
	value string
}

// Graph is the data structure representing a graph of Nodes
type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
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
