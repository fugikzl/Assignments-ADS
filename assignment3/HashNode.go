package assignment3

type HashNode struct {
	Key   StringableHashable //in our sample id of person
	Value Stringable
	Next  *HashNode
}

func NewHashNode(key StringableHashable, value Stringable) HashNode {
	newNode := HashNode{
		Key:   key,
		Value: value,
		Next:  nil,
	}

	return newNode
}

func (node *HashNode) ToString() string {
	return node.Key.String() + " : " + node.Value.String()
}
