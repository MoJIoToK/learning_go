# Homework 17

## Task

### Task 17.3.1

Напишите программу, аналогичную той, что мы только что написали, однако она должна использовать уже не 1000 горутин, а
только 10.

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Шаг наращивания счётчика
const STEP_1 int64 = 1

// Конечное значение счетчика
const endCounterValue int64 = 1000

func main() {

	var counter int64 = 0
	var wg sync.WaitGroup
	increment := func() {
		defer wg.Done()
		atomic.AddInt64(&counter, STEP_1)
	}
	// Не всегда вычисление этой переменной будет приводить к верному 
	// результату в счётчике, но для правильных значений 
	// и для удобства - можно
	var iterationCount int = int(endCounterValue / STEP_1)
	for i := 1; i <= iterationCount; i++ {
		wg.Add(1)
		go increment()
	}
	// Ожидаем поступления сигнала
	wg.Wait()
	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
```

### Task 17.3.3

Перепишите приведённый выше пример со счётчиком из основного текста, но вместо примитивов из пакета atomic используйте
условную переменную и попробуйте реализовать динамическую проверку достижения конечного значения счётчиком.

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const STEP_1 int64 = 1
const ITERATION_AMOUNT int = 1000

func main() {
	var counter int64 = 0
	var c = sync.NewCond(&sync.Mutex{})
	increment := func(i int) {
		atomic.AddInt64(&counter, STEP_1)
		if i == ITERATION_AMOUNT {
			c.Signal()
		}
	}
	for i := 1; i <= ITERATION_AMOUNT; i++ {
		go increment(i)
	}
	c.L.Lock()
	c.Wait()
	c.L.Unlock()
	fmt.Println(counter)
}
```

### Task 17.6.1

Напишите код, в котором имеются два канала сообщений из целых чисел, так, чтобы приём сообщений из них никогда не
приводил к блокировке и чтобы вероятность приёма сообщения из первого канала была выше в 2 раза, чем из второго.
*Если хотите, можете написать код, который бы демонстрировал это соотношение.

### Task 17.6.2

Напишите код, в котором имеются два канала сообщений из целых чисел так, чтобы приём сообщений всегда приводил к
блокировке. Приёмом сообщений из обоих каналов будет заниматься главная горутина. Сделайте так, чтобы во время такого
«бесконечного ожидания» сообщений выполнялась фоновая работа в виде вывода текущего времени в консоль.

### Task 17.6.3
Напишите программу, которая делает следующее: одна горутина по порядку отсылает числа от 1 до 100 в канал, вторая
горутина их принимает в правильном порядке и печатает на экран (в консоль).

## Solution

### Solution 17.3.1

[Решение находится здесь!](https://github.com/MoJIoToK/learning_go/blob/master/module17/homework_17.3.1.go)

### Solution 17.3.3

[Решение находится здесь!](https://github.com/MoJIoToK/learning_go/blob/master/module17/homework_17.3.3.go)

### Solution 17.6.1

[Решение находится здесь!](https://github.com/MoJIoToK/learning_go/blob/master/module17/task17.6.1.go)

### Solution 17.6.2

[Решение находится здесь!](https://github.com/MoJIoToK/learning_go/blob/master/module17/task17.6.2.go)

### Solution 17.6.3

[Решение находится здесь!](https://github.com/MoJIoToK/learning_go/blob/master/module17/task17.6.2.go)

## Improvements
