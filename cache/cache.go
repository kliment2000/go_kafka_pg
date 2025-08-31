package cache

import (
	"container/heap"
	"time"

	"github.com/kliment2000/go_kafka_pg/database"
)

type LRUCache struct {
	capacity int
	items    map[string]*Item
	pq       PriorityQueue
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		capacity: cap,
		items:    make(map[string]*Item),
		pq:       make(PriorityQueue, 0, cap),
	}
}

var Cache = NewLRUCache(100)

func (c *LRUCache) Get(key string) (interface{}, bool) {
	if item, ok := c.items[key]; ok {
		c.pq.update(item, item.value, time.Now().UnixNano())
		return item.value, true
	}
	return "", false
}

func (c *LRUCache) Put(key string, value interface{}) {
	if item, ok := c.items[key]; ok {
		c.pq.update(item, value, time.Now().UnixNano())
		return
	}

	if len(c.items) >= c.capacity {
		oldest := heap.Pop(&c.pq).(*Item)
		delete(c.items, oldest.key)
	}

	item := &Item{
		key:      key,
		value:    value,
		priority: time.Now().UnixNano(),
	}
	heap.Push(&c.pq, item)
	c.items[key] = item
}

func (c *LRUCache) LoadFromDB() error {
	orders, err := database.GetLastOrders(100)

	if err != nil {
		return err
	}

	for _, o := range orders {
		var jsonData interface{}
		if err := o.UnmarshalData(&jsonData); err != nil {
			return err
		}
		c.Put(o.OrderUID, jsonData)
	}

	return nil
}
