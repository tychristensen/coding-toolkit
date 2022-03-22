package graphs

import "sync"

// node is the interface representing a single node
type node interface {
	GetName() string
	String() string
	IsLessThan(n node) bool
}

// NodeBase is a struct that has the basic fields
// a node should have
type NodeBase struct {
	node
	lock sync.RWMutex
	Name string
}

func (n *NodeBase) GetName() string { return n.Name }
