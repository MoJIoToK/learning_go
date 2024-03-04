package main

import (
	"fmt"
	"sync"
)

const step int = 1
const interationAmount int = 1000

func main() {
	var counter int = 0
	var mu = sync.RWMutex{}
	var c = sync.NewCond(&mu)
	increment := func() {
		c.L.Lock()
		counter += step
		fmt.Println(counter)
		c.L.Unlock()

		//for {
		//if counter == interationAmount {
		//	c.Signal()
		//	}
		//}
	}
	for i := 1; i <= interationAmount; i++ {
		go increment()
	}

	go func() {
		for {
			if counter == interationAmount {
				c.Signal()
			}
		}
	}()

	c.L.Lock()
	c.Wait()
	c.L.Unlock()
	fmt.Println(counter)
}
