package assignment2

type Node struct {
	Previous *Node
	Next     *Node
	Element  interface{}
}

type MyLinkedList struct {
	Head     *Node //pointer to Node instance
	Tail     *Node //pointer to Node instance
	ListSize int
}

func NewMyLinkedList() *MyLinkedList {
	list := MyLinkedList{
		Head:     nil,
		Tail:     nil,
		ListSize: 0,
	}
	return &list
}

func (list *MyLinkedList) Add(item interface{}) {
	NewNode := &Node{ //empty node
		Element:  item,
		Previous: nil,
		Next:     nil,
	}

	if list.Head == nil { //if head is empty setting head and tail as node
		list.Head = NewNode
		list.Tail = NewNode
	} else { //if not, tail beocme previous and next become out given item
		NewNode.Previous = list.Tail
		list.Tail.Next = NewNode
		list.Tail = NewNode
	}
	list.ListSize++
}

func (list *MyLinkedList) Get(index int) interface{} {
	if index >= 0 && index < list.ListSize { //check if we have that index node
		var CurrentNode *Node
		if index < list.ListSize/2 { //dividing by two to determinw from which part is closer
			CurrentNode = list.Head
			for i := 0; i < index; i++ {
				CurrentNode = CurrentNode.Next
			}
		} else {
			CurrentNode = list.Tail
			for i := list.ListSize - 1; i > index; i-- {
				CurrentNode = CurrentNode.Previous
			}
		}

		return CurrentNode.Element
	}

	return nil //if not return niil
}

func (list *MyLinkedList) Size() int {
	return list.ListSize
}

func (list *MyLinkedList) Remove(index int) bool {
	if index < 0 || index >= list.ListSize {
		return false
	}

	var currentNode *Node
	if index < list.ListSize/2 { //dividing by two to determinw from which part is closer
		currentNode = list.Head
		for i := 0; i < index; i++ {
			currentNode = currentNode.Next
		}
	} else {
		currentNode = list.Tail
		for i := list.ListSize - 1; i > index; i-- {
			currentNode = currentNode.Previous
		}
	}

	if currentNode.Previous != nil { //if head
		currentNode.Previous.Next = currentNode.Next
	} else {
		list.Head = currentNode.Next //next node become head
	}

	if currentNode.Next != nil { //if tail
		currentNode.Next.Previous = currentNode.Previous // {s,b,c,d} : remove b -> {s} {c,d}  c.previous is s
	} else {
		list.Tail = currentNode.Previous // previous of current is tail now
	}

	list.ListSize--
	return true
}

func (list *MyLinkedList) Contains(item interface{}) bool {
	currentNode := list.Head

	for currentNode != nil {
		if currentNode.Element == item {
			return true
		}
		currentNode = currentNode.Next
	}

	return false
}

func (list *MyLinkedList) Clear() {
	list.Head = nil   //set to nil
	list.Tail = nil   //set to nil
	list.ListSize = 0 //set to 0
}

func (list *MyLinkedList) IndexOf(item interface{}) int {
	currentNode := list.Head //starting form head
	index := 0               //by defult is zero

	for currentNode != nil {
		if currentNode.Element == item {
			return index
		}
		index++
		currentNode = currentNode.Next
	}

	return -1
}

func (list *MyLinkedList) LastIndexOf(item interface{}) int {
	currentNode := list.Tail
	index := list.ListSize - 1

	for currentNode != nil {
		if currentNode.Element == item {
			return index
		}
		index--
		currentNode = currentNode.Previous
	}

	return -1
}

// type ListInterface interface {
// 	Size() int c
// 	Contains(interface{}) bool
// 	Add(interface{})
// 	Remove(interface{}) bool
// 	Clear()
// 	Get(int) interface{}
// 	IndexOf(interface{}) int
// 	LastIndexOf(interface{}) int
// }
