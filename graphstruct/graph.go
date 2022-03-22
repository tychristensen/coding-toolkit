package graphs

import "sync"

// Graph is the data structure representing a graph of Nodes
// Do not reference the struct directly use the NewGraph()
// function instead
type Graph struct {
	nodes []*NodeBase
	edges map[string][]*NodeBase
	lock  sync.RWMutex
}
type Add interface {
	Edge()
	Node()
}

func NewGraph() *Graph {
	return &Graph{edges: make(map[string][]*NodeBase)}
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(n *NodeBase) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(n1, n2 *NodeBase) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.edges[(*n1).GetName()] = append(g.edges[(*n1).GetName()], n2)
	g.edges[(*n2).GetName()] = append(g.edges[(*n2).GetName()], n1)
}

// String returns a string version of a graph
func (g *Graph) String() (str string) {
	for _, node := range g.nodes {
		str += node.GetName() + "\n"
		for _, edge := range g.edges[node.GetName()] {
			str += "\t=>" + edge.String() + "\n"
		}
		str += "\n    "
	}

	return str
}
