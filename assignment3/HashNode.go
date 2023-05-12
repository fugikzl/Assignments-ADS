package assignment3

import (
	"Assignments-ADS/interfaces"
)

type HashNode struct {
	Key   interfaces.StringableHashable //in our sample id of person
	Value interfaces.Stringable
	Next  *HashNode
}

func NewHashNode(key interfaces.StringableHashable, value interfaces.Stringable) HashNode {
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
