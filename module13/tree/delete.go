package main

func (tree *Tree) Delete(value int) bool {

	if !tree.Find(value) || tree.root == nil {
		return false
	}

	if tree.root.Value == value {
		temp := &Node{nil, nil, 0}
		temp.Left = tree.root
		r := delete(tree.root, temp, value)
		tree.root = temp.Left
		return r
	}
	return delete(tree.root.Left, tree.root, value) || delete(tree.root.Right, tree.root, value)
}

func delete(node *Node, parent *Node, value int) bool {
	switch {
	case node.Value == value:
		if node.Left != nil && node.Right != nil {
			node.Value = minValue(node.Right)
			return delete(node.Right, node, node.Value)
		}
		link(parent, node)
		return true
	case node.Value > value:
		if node.Left == nil {
			return false
		}
		return delete(node.Left, node, value)
	case node.Value < value:
		if node.Right == nil {
			return false
		}
		return delete(node.Right, node, value)
	}
	return false
}

func minValue(node *Node) int {
	if node.Left == nil {
		return node.Value
	}
	return minValue(node.Right)
}

func link(parent, node *Node) {
	switch {
	case parent.Left == node:
		if node.Left != nil {
			parent.Left = node.Left
		} else {
			parent.Left = node.Right
		}
	case parent.Right == node:
		if node.Left != nil {
			parent.Right = node.Left
		} else {
			parent.Right = node.Right
		}
	}
}
