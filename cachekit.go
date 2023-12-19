// Package cachekit provides a thread-safe cache implementation with support for generic key-value types.
package cachekit

import (
	"fmt"
	"sync"
)

// CacheKit is a thread-safe cache implementation with support for generic key-value types.
type CacheKit[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// New creates a new instance of CacheKit.
func New[K comparable, V any]() *CacheKit[K, V] {
	return &CacheKit[K, V]{
		data: make(map[K]V),
	}
}

// Get retrieves the value associated with the given key from the cache.
// It returns the value and a boolean indicating whether the key was found in the cache.
func (c *CacheKit[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

// Set sets the value associated with the given key in the cache.
func (c *CacheKit[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Update updates the value associated with the given key in the cache.
// If the key does not exist in the cache, an error is returned.
func (c *CacheKit[K, V]) Update(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.data[key]; !ok {
		return fmt.Errorf("key %v does not exist", key)
	}
	c.data[key] = value
	return nil
}

// Delete removes the value associated with the given key from the cache.
func (c *CacheKit[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Len returns the number of items in the cache.
func (c *CacheKit[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}

// Clear removes all items from the cache.
func (c *CacheKit[K, V]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[K]V)
}

// Keys returns a slice containing all the keys in the cache.
func (c *CacheKit[K, V]) Keys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()
	keys := make([]K, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice containing all the values in the cache.
func (c *CacheKit[K, V]) Values() []V {
	c.mu.RLock()
	defer c.mu.RUnlock()
	values := make([]V, 0, len(c.data))
	for _, v := range c.data {
		values = append(values, v)
	}
	return values
}

// Items returns a map containing all the key-value pairs in the cache.
func (c *CacheKit[K, V]) Items() map[K]V {
	c.mu.RLock()
	defer c.mu.RUnlock()
	items := make(map[K]V, len(c.data))
	for k, v := range c.data {
		items[k] = v
	}
	return items
}

// Has checks if the cache contains the given key.
func (c *CacheKit[K, V]) Has(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.data[key]
	return ok
}
