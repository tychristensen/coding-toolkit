package graphs

type numNode struct {
	NodeBase
	Value uint64
}

func (n1 *numNode) IsLessThan(n2 *numNode) bool {
	return n1.Value < n2.Value
}
