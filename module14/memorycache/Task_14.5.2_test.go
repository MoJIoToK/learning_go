package memorycache

import (
	"testing"
	"time"
)

const (
	testKey      string = "inMemoryCache:test"
	testKeyEmpty string = "inMemoryCache:empty"
	testValue    string = "Hello Test"
)

// AppCache init new cache.
var AppCache = NewInMemoryCache(5*time.Second, 10*time.Second)

//var AppCacheGC = NewInMemoryCache(5*time.Second, 1*time.Second)

// TestGet get cache by its key.
func TestGet(t *testing.T) {
	AppCache.Set(testKey, testValue, 2*time.Second)

	value, ok := AppCache.Get(testKey)

	if value != testValue {
		t.Error("Error: ", "The received value: do not correspond to the expectation:", value, testValue)
	}

	if ok != true {
		t.Error("Error: ", "Could not get cache")
	}

	value, ok = AppCache.Get(testKeyEmpty)

	if value != nil || ok != false {
		t.Error("Error: ", "Value does not exist and must be empty", value)
	}

}

// TestDelete delete cache by key
func TestDelete(t *testing.T) {

	AppCache.Set(testKey, testValue, 1*time.Minute)

	error := AppCache.Delete(testKey)

	if error != nil {
		t.Error("Error: ", "Cache delete failed")
	}

	value, found := AppCache.Get(testKey)

	if found {
		t.Error("Error: ", "Should not be found because it was deleted")
	}

	if value != nil {
		t.Error("Error: ", "Value is not nil:", value)
	}

	// repeat deletion of an existing cache
	error = AppCache.Delete(testKeyEmpty)

	if error == nil {
		t.Error("Error: ", "An empty cache should return an error")
	}

}
