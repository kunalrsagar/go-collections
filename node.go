package collections

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value, Next: nil}
}
