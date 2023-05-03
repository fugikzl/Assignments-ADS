package assignment3

import "fmt"

type MyHashTable struct {
	chainArray []*HashNode
	m          int
	size       int
}

func NewMyHashTable(m int) *MyHashTable {
	table := MyHashTable{
		m:          m,
		size:       0,
		chainArray: make([]*HashNode, m),
	}

	return &table
}

func (t *MyHashTable) hash(key StringableHashable) int {
	return key.HashCode() % t.m
}

func (t *MyHashTable) Put(key StringableHashable, value Stringable) {
	h := t.hash(key)
	node := HashNode{
		Key:   key,
		Value: value,
		Next:  t.chainArray[h],
	}
	t.chainArray[h] = &node
}

func (t *MyHashTable) Get(key StringableHashable) (StringableHashable, Stringable) {
	h := t.hash(key)
	for node := t.chainArray[h]; node != nil; node = node.Next {
		if key.Equals(node.Key) {
			return node.Key, node.Value
		}
	}
	return nil, nil
}

func (t *MyHashTable) PrintChainArrayElement(i int) {
	for node := t.chainArray[i]; node != nil; node = node.Next {
		fmt.Print(node, " ")
	}
	fmt.Print("\n")
}
