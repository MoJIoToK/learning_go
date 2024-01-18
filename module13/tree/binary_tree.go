package main

// region Структуры
type Node struct {
	Left  *Node
	Right *Node
	Value int
}

type Tree struct {
	root *Node
	size int
}

//endregion

// region Конструктор
func NewTree() *Tree {
	tree := new(Tree)
	tree.size = 0
	return tree
}

//endregion

//region Методы

// Size returns size of binary tree
func (tree *Tree) Size() int {
	return tree.size
}

// Root returns root of binary tree
func (tree *Tree) Root() *Node {
	return tree.root
}

//endregion
