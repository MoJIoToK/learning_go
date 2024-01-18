package main

// Public method Insert adds node in tree.
// If subroot does not exist, then it is created with the default values of the left and right child.
func (tree *Tree) Insert(value int) {
	if tree.root == nil {
		tree.root = &Node{nil, nil, value}
	}
	tree.size++
	tree.root.insert(&Node{nil, nil, value})
}

// Insert function adds node in binary tree. If node is greater than subroot, then node add in right.
// If node is less than subroot, then node add in left.
func (root *Node) insert(new_node *Node) {
	switch {
	case new_node.Value > root.Value:
		if root.Right == nil {
			root.Right = new_node
		} else {
			root.Right.insert(new_node)
		}
	case new_node.Value < root.Value:
		if root.Left == nil {
			root.Left = new_node
		} else {
			root.Left.insert(new_node)
		}
	}
}
