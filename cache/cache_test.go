package cache

import (
	"testing"
	"time"
)

func TestLRUCachePutAndGet(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", "1")
	cache.Put("b", "2")

	if v, ok := cache.Get("a"); !ok || v != "1" {
		t.Errorf("expected a=1, got %v, ok=%v", v, ok)
	}

	if v, ok := cache.Get("b"); !ok || v != "2" {
		t.Errorf("expected b=2, got %v, ok=%v", v, ok)
	}
}

func TestLRUCacheUpdateExisting(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", "1")
	cache.Put("a", "updated")

	if v, ok := cache.Get("a"); !ok || v != "updated" {
		t.Errorf("expected updated value for a, got %v", v)
	}
}

func TestLRUCacheCapacityOne(t *testing.T) {
	cache := NewLRUCache(1)

	cache.Put("a", "1")
	cache.Put("b", "2")

	if _, ok := cache.Get("a"); ok {
		t.Error("expected a to be evicted with capacity=1")
	}
	if v, ok := cache.Get("b"); !ok || v != "2" {
		t.Errorf("expected b=2, got %v", v)
	}
}

func TestLRUCacheTouchUpdatesPriority(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", "1")
	time.Sleep(1 * time.Millisecond)
	cache.Put("b", "2")
	time.Sleep(1 * time.Millisecond)

	cache.Get("a")

	cache.Put("c", "3")

	if _, ok := cache.Get("b"); ok {
		t.Error("expected b to be evicted after c insert")
	}
	if _, ok := cache.Get("a"); !ok {
		t.Error("expected a to remain after using it")
	}
	if _, ok := cache.Get("c"); !ok {
		t.Error("expected c to be present")
	}
}
