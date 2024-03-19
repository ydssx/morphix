package cache

import (
	"time"
)

type Cache interface {
	// Get 从缓存中获取指定key的值,并反序列化到result中
	Get(key string, result interface{}) error
	// Set 将指定的key/value对设置到缓存中,并设置过期时间
	Set(key string, value interface{}, expire time.Duration) error
	// Delete 从缓存中删除指定key的值
	Delete(key string) error
	// Clear 清空缓存中的所有键值对
	Clear() error
}
