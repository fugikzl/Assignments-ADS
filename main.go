package main

import (
	"Assignments-ADS/assignment4"
	"fmt"
)

func main() {
	BST := assignment4.NewBinarySearchTree()

	BST.Put(8, "a")
	BST.Put(9, "a")
	BST.Put(3, "a")
	BST.Put(4, "a")
	BST.Put(5, "a")
	BST.Put(6, "a")
	BST.Put(10, "a")

	fmt.Print(BST.GetRoot().Left.Right.Right.Get())
}
