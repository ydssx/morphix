package cache

import (
	"bytes"
	"encoding/gob"
	"sync"
	"time"
)

type MemoryCache struct {
	sync.Mutex
	expirations sync.Map
	data        map[string]interface{}
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{data: make(map[string]interface{})}
}

func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()
	value, ok := c.data[key]
	return value, ok
}

func (c *MemoryCache) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.data, key)
}

func (c *MemoryCache) Clear() {
	c.Lock()
	defer c.Unlock()
	c.data = make(map[string]interface{})
}

func (c *MemoryCache) Set(key string, value interface{}, expiration time.Duration) {
	c.Lock()
	defer c.Unlock()

	c.data[key] = value

	// 如果设置了过期时间，则开启定时器，在到达过期时间时自动删除该键值对
	if expiration > 0 {
		expirationTime := time.Now().Add(expiration)
		c.expirations.Store(key, expirationTime)
		go func() {
			timer := time.NewTimer(expiration)
			defer timer.Stop()

			// 在定时器到期之前如果该键值对被更新或删除了，则不进行删除操作
			select {
			case <-timer.C:
				c.Lock()
				defer c.Unlock()

				// 原子操作检查键值对是否过期
				if expire, ok := c.expirations.Load(key); ok && expire.(time.Time).Equal(expirationTime) {
					delete(c.data, key)
					c.expirations.Delete(key)
				}
			}
		}()
	} else {
		c.expirations.Delete(key)
	}
}

// 使用 gob 序列化和反序列化实现深拷贝
func (c *MemoryCache) deepCopy(value interface{}) interface{} {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	decoder := gob.NewDecoder(buffer)

	// 编码 value
	encoder.Encode(value)

	// 解码 value 的副本
	var copiedValue interface{}
	decoder.Decode(&copiedValue)

	return copiedValue
}
