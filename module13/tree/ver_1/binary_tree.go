package main

import (
	"errors"
	"fmt"
	"log"
	"module13/tree/ver_1/node"
	"strconv"
)

type Tree struct {
	Root *node.Node
}

func (t *Tree) Insert(value int, data string) error {
	if t.Root == nil {
		t.Root = &node.Node{Value: value, Data: data}
		return nil
	}
	return t.Root.Insert(value, data)
}

func (t *Tree) Find(s int) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(s)
}

func (t *Tree) Delete(s int) error {
	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}
	fakeParent := &node.Node{Right: t.Root}
	err := t.Root.Delete(s, fakeParent)
	if err != nil {
		return err
	}
	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

func (t *Tree) Traverse(n *node.Node, f func(*node.Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}

	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}

	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *node.Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	s := 5
	fmt.Print("Find node ", s, ": ")
	d, found := tree.Find(s)
	if !found {
		log.Fatal("Cannot find " + strconv.Itoa(s) + "")
	}
	fmt.Println("Found " + strconv.Itoa(s) + ": " + d + "")

	err := tree.Delete(s)
	if err != nil {
		log.Fatal("Error deleting "+strconv.Itoa(s)+": ", err)
	}
	fmt.Print("After deleting " + strconv.Itoa(s) + ": ")
	tree.Traverse(tree.Root, func(n *node.Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

}
