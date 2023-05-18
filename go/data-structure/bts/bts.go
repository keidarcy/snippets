package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add ad node to the tree
// the key to add should not be already in the key
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	}

	if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(52)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(56)
	tree.Insert(88)
	tree.Insert(203)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(276)

	found := tree.Search(88)
	fmt.Printf("found: %v\n", found)
}
