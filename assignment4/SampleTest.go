package assignment4

import "fmt"

func TestBST() {
	BST := NewBinarySearchTree()

	BST.Put(8, "a")
	BST.Put(9, "b")
	BST.Put(3, "c")
	BST.Put(4, "d")
	BST.Put(5, "e")
	BST.Put(6, "f")
	BST.Put(10, "g")

	a, b := BST.GetRoot().Right.Right.Get()
	if a == 10 && b == "g" {
		fmt.Print("\nPut Test N1 Is OK")
	} else {
		fmt.Print("\nPut Test N1 Is NOT OK")
	}

	a, b = BST.GetRoot().Left.Right.Right.Get()
	if a == 5 && b == "e" {
		fmt.Print("\nPut Test N2 Is OK")
	} else {
		fmt.Print("\nPut Test N2 Is NOT OK")
	}

	b = BST.Get(4)
	if b == "d" {
		fmt.Print("\nGet Test N1 Is OK")
	} else {
		fmt.Print("\nGet Test N1 Is NOT OK")
	}

	b = BST.Get(3)
	if b == "c" {
		fmt.Print("\nGet Test N2 Is OK")
	} else {
		fmt.Print("\nGet Test N2 Is NOT OK")
	}
}
