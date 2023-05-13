package assignment4

type BinarySearchTree struct {
	root *Node
	size uint64
}

func (BST *BinarySearchTree) GetRoot() *Node {
	return BST.root
}

func NewBinarySearchTree() *BinarySearchTree {
	BST := BinarySearchTree{
		root: nil,
		size: 0,
	}

	return &BST
}

func (BST *BinarySearchTree) put(k int, v string, n *Node) *Node {
	if n == nil {
		return &Node{
			key:   k,
			value: v,
		}
	}
	if k < n.key {
		n.Left = BST.put(k, v, n.Left)
	} else if k > n.key {
		n.Right = BST.put(k, v, n.Right)
	} else {
		n.value = v
	}

	return n
}

func (BST *BinarySearchTree) Put(k int, v string) {
	BST.root = BST.put(k, v, BST.root)
	BST.size++
}

func (BST *BinarySearchTree) get(k int, n *Node) *Node {
	if n.key == k {
		return n
	}

	if n.key < k {
		return BST.get(k, n.Right)
	}

	if n.key > k {
		return BST.get(k, n.Left)
	}

	return nil

}

func (BST *BinarySearchTree) Get(k int) string {
	n := BST.get(k, BST.root)
	if n == nil {
		return "No foind"
	}

	return n.value
}

func (BST *BinarySearchTree) Delete(k int) {

}

// I'm communised code
// Freakin go, code absolutely no readable

// public void put(K key, V value) {
// 	root = put(root, key, value);
// 	size++;
// }

// private Node put(Node x, K key, V value) {
// 	if (x == null) {
// 		return new Node(key, value);
// 	}
// 	if (key.compareTo(x.key) < 0) {
// 		x.left = put(x.left, key, value);
// 	} else if (key.compareTo(x.key) > 0) {
// 		x.right = put(x.right, key, value);
// 	} else {
// 		x.value = value;
// 	}
// 	return x;
// }
