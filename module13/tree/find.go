package main

// Public method returns boolean value.
func (tree *Tree) Find(value int) bool {
	tree.size--
	return find(tree.root, value)
}

// Find returns true, if target value exists in tree.
// The function calls itself recursively depending on whether the target value is greater or less than subroot.
// If target value is greater than subroot, then function called recursively calls the function with the right child.
func find(root *Node, value int) bool {
	switch {
	case root != nil:
		if value == root.Value {
			return true
		} else if value > root.Value {
			return find(root.Right, value)
		} else {
			return find(root.Left, value)
		}
	}
	return false
}
