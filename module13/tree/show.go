package main

import "fmt"

// Show displays to the console nodes in ascending order.
func show(root *Node) {
	if root != nil {
		show(root.Left)
		fmt.Println(root.Value)
		show(root.Right)
	}
}

// ShowPreOrder displays to the console root - left sub-tree - right sub-tree.
func showPreOrder(root *Node) {
	if root != nil {
		fmt.Println(root.Value)
		show(root.Left)
		show(root.Right)
	}

}
