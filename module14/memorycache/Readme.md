# Task 14.5.2

## Task
Используя мапу, реализуйте тип InMemoryCache, который позволит хранить значения в 
течение какого-то определённого времени (InMemoryCache должен реализовывать Cache interface):
```go
package main
import "time"

var _ Cache = InMemoryCache{} // это трюк для проверки типа: до тех пор пока InMemoryCache не будет реализовывать интерфейс Cache, программа не запустится

type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type InMemoryCache struct {
      //Code here
}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		// Code here
	}
}
```

## Solution

[Решение смотри здесь](https://github.com/MoJIoToK/learning_go/tree/master/module14/memorycache)

## References

[Пишем простой менеджер кеша в памяти на Go](https://habr.com/ru/articles/359078/)