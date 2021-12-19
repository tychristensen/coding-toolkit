package main

func main() {
	graph := Node{}
	a := &Node{value: "a"}
	b := &Node{value: "b"}
	graph.AddNode(a)
	graph.AddNode(b)
	graph.AddEdge(a, b)
}
