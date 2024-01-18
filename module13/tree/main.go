package main

import "fmt"

func main() {
	test := NewTree()

	test.Insert(4)
	test.Insert(1)
	test.Insert(11)
	test.Insert(22)
	test.Insert(699)
	test.Insert(3)

	//fmt.Println(test.Size())

	show(test.root)
	fmt.Println("\nshowPreOrder:")
	showPreOrder(test.root)
	fmt.Println("\nDelete 22")
	test.Delete(22)
	fmt.Println("\nshowPreOrder after delete")
	showPreOrder(test.root)
	fmt.Println("\nFind 4")
	fmt.Println(test.Find(4))
	fmt.Println("\nFind 111")
	fmt.Println(test.Find(111))

}
