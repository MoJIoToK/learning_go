package main

import "fmt"

// ErrIndexOutOfBounds - описание ошибки при обращении к элементу с неверным индексом
var ErrIndexOutOfBounds = fmt.Errorf("Ошибка доступа за границы структуры")

// Node - узел списка
type Node struct {
	Value int
	next  *Node
}

// List - связный список элементов, хранящих целые числа
type List struct {
	size int
	head *Node
}

// Find - поиск узла по индексу
func (linkedList List) Find(index int) (*Node, error) {
	if index < 0 || index >= linkedList.size {
		return nil, ErrIndexOutOfBounds
	}
	var node *Node = linkedList.head
	for i := 1; i <= index; i++ {
		node = node.next
	}
	return node, nil
}

func (l *List) add(val int) {
	newNode := &Node{Value: val}
	if l.head == nil {
		l.head = newNode
		l.size++
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
	l.size++
}

func printList(l *List) {
	cur := l.head
	for cur != nil {
		fmt.Println(cur.Value)
		cur = cur.next
	}
}

func main() {
	list := &List{}
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)
	printList(list)
	println(list.size)
	fmt.Println(list.Find(1))
}
