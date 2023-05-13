package assignment4

type Node struct {
	key         int
	value       string
	Left, Right *Node
}

func (n *Node) Get() (int, string) {
	return n.key, n.value
}

func NewNode(k int, v string) *Node {
	node := Node{
		key:   k,
		value: v,
		Left:  nil,
		Right: nil,
	}

	return &node
}
