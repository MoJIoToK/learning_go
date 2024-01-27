package main

import (
	"errors"
	"sync"
	"time"
)

var _ Cache = &InMemoryCache{} // это трюк для проверки типа:
// до тех пор пока InMemoryCache не будет реализовывать интерфейс Cache, программа не запустится

// CacheEntry структура описывающая элементы кеша.
type CacheEntry struct {
	//SettledAt - время создания элемента.
	SettledAt time.Time
	//Value - значение кеша, т.к. оно может быть любого типа, то тип interface{}
	Value interface{}
	//Expiration - время истечения (в UnixNano). По нему проверяеться актуальность кеша.
	Expiration int64
}

type Cache interface {
	Set(key string, value interface{}, duration time.Duration)
	Get(key string) (interface{}, bool)
}

// InMemoryCache контейнер-хранилище.
type InMemoryCache struct {
	sync.RWMutex
	//DefaultExpiration - продолжительность хранения кеша по-умолчанию. Можно переопределить.
	//Если установлено значение меньше или равно 0 - время жизни бесконечно.
	defaultExpiration time.Duration
	//CleanupInterval - интервал, через который запускается механизм удаления кеша (Garbage Collector/GC).
	//Если установлено значение меньше или равно 0 - удаление просроченного кеша не происходит.
	cleanupInterval time.Duration
	//Items - элементы хранящиеся в кеше. Хранение в формате ключ-значение. Где значением является структура CacheEntry
	items map[string]CacheEntry
}

// NewInMemoryCache - конструктор. Инициализация хранилища.
func NewInMemoryCache(defaultExpiration, cleanupInterval time.Duration) *InMemoryCache {
	items := make(map[string]CacheEntry)

	inMemoryCache := InMemoryCache{
		items:             items,
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}

	//Если интервал очистки больше 0, то запускается процесс очистки кеша GC.
	if cleanupInterval > 0 {
		inMemoryCache.StartGC()
	}

	return &inMemoryCache
}

// Set - добавляет новый элемент в кеш или заменяет существующий.
// В качестве аргументов принимает: ключ-идентификатор в виде строки key,
// значение value и продолжительность жизни кеша duration.
func (imc *InMemoryCache) Set(key string, value interface{}, duration time.Duration) {

	//expiration целое число.
	var expiration int64

	if duration == 0 {
		duration = imc.defaultExpiration
	}
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	imc.Lock()

	defer imc.Unlock()

	imc.items[key] = CacheEntry{
		Value:      value,
		Expiration: expiration,
		SettledAt:  time.Now(),
	}

}

// Get - возвращает значение (или nil) и второй параметр типа bool. Если кеш удален или устарел возвращается nil, false.
func (imc *InMemoryCache) Get(key string) (interface{}, bool) {
	imc.RLock()
	defer imc.RUnlock()

	item, ok := imc.items[key]

	if !ok {
		return nil, false
	}

	//Проверка на время истечения.
	if item.Expiration > 0 {
		// Проверка не устарел ли кеш в момент запроса
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}

	}

	return item.Value, true
}

// Delete удаляет элемент кеша по ключу. Возвращает ошибку, если такого ключа нет.
func (imc *InMemoryCache) Delete(key string) error {
	imc.Lock()

	defer imc.Unlock()

	if _, ok := imc.items[key]; !ok {
		return errors.New("Key not found!")
	}

	delete(imc.items, key)

	return nil
}

func (imc *InMemoryCache) StartGC() {
	go imc.GC()
}

func (imc *InMemoryCache) GC() {
	for {
		<-time.After(imc.cleanupInterval)

		if imc.items == nil {
			return
		}

		if keys := imc.expiredKeys(); len(keys) != 0 {
			imc.clearItems(keys)
		}

	}

}

func (imc *InMemoryCache) expiredKeys() (keys []string) {
	imc.RLock()

	defer imc.RUnlock()

	for key, val := range imc.items {
		if time.Now().UnixNano() > val.Expiration && val.Expiration > 0 {
			keys = append(keys, key)
		}
	}
	return
}

func (imc *InMemoryCache) clearItems(keys []string) {
	imc.Lock()

	defer imc.Unlock()

	for _, key := range keys {
		delete(imc.items, key)
	}

}

//Используя мапу, реализуйте тип InMemoryCache, который позволит хранить значения в течение
//какого-то определённого времени (InMemoryCache должен реализовывать Cache interface):
