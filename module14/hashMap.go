package main

import "fmt"

const (
	HASH_MOD = 1000
	MAP_SIZE = 10
)

type hashElement struct {
	key   string
	value string
	next  *hashElement
}

type HashMap struct {
	size     int
	elements []*hashElement
}

func NewHashMap(size int) *HashMap {
	return &HashMap{size: size, elements: make([]*hashElement, size)}
}

func (hm *HashMap) hash(key string) uint64 {
	//hash := 0
	//for _, char := range key {
	//	hash += int(char)
	//}
	//return hash % hm.size
	hash := uint64(0)
	for _, char := range key {
		hash = hash*31 + uint64(char)
	}
	return hash % HASH_MOD % MAP_SIZE
}

func (hm *HashMap) Set(key, value string) {
	index := hm.hash(key)
	newElement := &hashElement{key: key, value: value}
	if hm.elements[index] == nil {
		hm.elements[index] = newElement
	} else {
		current := hm.elements[index]
		for current.next != nil {
			current = current.next
		}
		current.next = newElement
	}
}

func (hm *HashMap) Get(key string) (string, bool) {
	index := hm.hash(key)
	current := hm.elements[index]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return "", false
}

func (hm *HashMap) Delete(key string) {
	_, ok := hm.Get(key)
	index := hm.hash(key)
	if ok {
		hm.elements[index] = nil
		fmt.Println("This key is deleted!")
	} else {
		fmt.Println("This key is not exist!")
	}
}

func main() {
	table := NewHashMap(MAP_SIZE)
	table.Set("First", "Alisa")
	table.Set("Second", "Bob")
	table.Set("Third", "Sam")
	fmt.Println(table)

	id, ok := table.Get("Second")
	if ok {
		fmt.Println(id)
	} else {
		fmt.Println("not found!")
	}

	table.Delete("Second")
	fmt.Println(table)

}

//func (h *hashmap) Delete(key string) {
//}
