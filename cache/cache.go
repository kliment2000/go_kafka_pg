package cache

import (
	"sync"

	"github.com/kliment2000/go_kafka_pg/db"
)

type CachedOrder struct {
	OrderUID string
	Data     interface{} // готовый распарсенный JSON
}

type orderCache struct {
	mu     sync.RWMutex
	orders map[string]CachedOrder
}

var Cache = &orderCache{
	orders: make(map[string]CachedOrder),
}

func (c *orderCache) SetOrder(order CachedOrder) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.orders[order.OrderUID] = order
}

func (c *orderCache) GetOrder(orderUID string) (CachedOrder, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	order, ok := c.orders[orderUID]
	return order, ok
}

func (c *orderCache) LoadFromDB() error {
	orders, err := db.GetAllOrders()

	if err != nil {
		return err
	}

	for _, o := range orders {
		var jsonData interface{}
		if err := o.UnmarshalData(&jsonData); err != nil {
			return err
		}
		c.SetOrder(CachedOrder{
			OrderUID: o.OrderUID,
			Data:     jsonData,
		})
	}

	return nil
}
