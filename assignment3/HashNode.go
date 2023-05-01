package assignment3

import (
	"Assignments-ADS/sampletypes"
	"strconv"
)

type HashNode struct {
	Key   int //in our sample id of person
	Value *sampletypes.Person
	Next  *HashNode
}

func NewHashNode(key int, value *sampletypes.Person) *HashNode {
	newNode := HashNode{
		Key:   key,
		Value: value,
		Next:  nil,
	}

	return &newNode
}

func (node *HashNode) ToString() string {
	return strconv.Itoa(node.Key) + node.Value.Name + node.Value.Job //idk is there any better way to convert string from int
}
