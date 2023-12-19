package cachekit

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheKit(t *testing.T) {
	c := New[string, int]()
	c.Set("hello", 1)
	c.Set("world", 2)
	c.Set("cachekit", 3)
	assert.Equal(t, 3, c.Len())

	testGet(t, c)
	testDelete(t, c)
	testClear(t, c)
}

func testGet(t *testing.T, c *CacheKit[string, int]) {
	val, ok := c.Get("hello")
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	val, ok = c.Get("world")
	assert.True(t, ok)
	assert.Equal(t, 2, val)
	val, ok = c.Get("cachekit")
	assert.True(t, ok)
	assert.Equal(t, 3, val)
}

func testDelete(t *testing.T, c *CacheKit[string, int]) {
	c.Delete("hello")
	assert.Equal(t, 2, c.Len())
	val, ok := c.Get("hello")
	assert.False(t, ok)
	assert.Equal(t, 0, val)
	val, ok = c.Get("world")
	assert.True(t, ok)
	assert.Equal(t, 2, val)
	val, ok = c.Get("cachekit")
	assert.True(t, ok)
	assert.Equal(t, 3, val)
}

func testClear(t *testing.T, c *CacheKit[string, int]) {
	c.Clear()
	assert.Equal(t, 0, c.Len())
	val, ok := c.Get("hello")
	assert.False(t, ok)
	assert.Equal(t, 0, val)
	val, ok = c.Get("world")
	assert.False(t, ok)
	assert.Equal(t, 0, val)
	val, ok = c.Get("cachekit")
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}

func TestCacheKitUpdate(t *testing.T) {
	c := New[string, int]()
	c.Set("hello", 1)
	c.Set("world", 2)
	c.Set("cachekit", 3)
	assert.Equal(t, 3, c.Len())

	testUpdate(t, c)
}

func testUpdate(t *testing.T, c *CacheKit[string, int]) {
	c.Update("hello", 4)
	assert.Equal(t, 3, c.Len())
	val, ok := c.Get("hello")
	assert.True(t, ok)
	assert.Equal(t, 4, val)
	val, ok = c.Get("world")
	assert.True(t, ok)
	assert.Equal(t, 2, val)
	val, ok = c.Get("cachekit")
	assert.True(t, ok)
	assert.Equal(t, 3, val)

	c.Update("foo", 5)
	assert.Equal(t, 3, c.Len())
	val, ok = c.Get("foo")
	assert.False(t, ok)
	assert.Equal(t, 0, val)
	val, ok = c.Get("hello")
	assert.True(t, ok)
	assert.Equal(t, 4, val)
	val, ok = c.Get("world")
	assert.True(t, ok)
	assert.Equal(t, 2, val)
	val, ok = c.Get("cachekit")
	assert.True(t, ok)
	assert.Equal(t, 3, val)
}

func TestCacheKitKeys(t *testing.T) {
	c := New[string, int]()
	c.Set("hello", 1)
	c.Set("world", 2)
	c.Set("cachekit", 3)
	assert.Equal(t, 3, c.Len())

	testKeys(t, c)
}

func testKeys(t *testing.T, c *CacheKit[string, int]) {
	keys := c.Keys()
	assert.Equal(t, 3, len(keys))
	assert.Contains(t, keys, "hello")
	assert.Contains(t, keys, "world")
	assert.Contains(t, keys, "cachekit")
}

func TestCacheKitConcurrency(t *testing.T) {
	c := New[string, int]()
	c.Set("hello", 1)
	c.Set("world", 2)
	c.Set("cachekit", 3)
	assert.Equal(t, 3, c.Len())

	testConcurrency(t, c)
}

func testConcurrency(t *testing.T, c *CacheKit[string, int]) {
	// Get
	go func() {
		val, ok := c.Get("hello")
		assert.True(t, ok)
		assert.Equal(t, 1, val)
	}()
	go func() {
		val, ok := c.Get("world")
		assert.True(t, ok)
		assert.Equal(t, 2, val)
	}()
	go func() {
		val, ok := c.Get("cachekit")
		assert.True(t, ok)
		assert.Equal(t, 3, val)
	}()

	// Set
	go func() {
		c.Set("foo", 4)
	}()
	go func() {
		c.Set("bar", 5)
	}()
	go func() {
		c.Set("baz", 6)
	}()

	// Update
	go func() {
		c.Update("hello", 7)
	}()
	go func() {
		c.Update("world", 8)
	}()
	go func() {
		c.Update("cachekit", 9)
	}()

	// Delete
	go func() {
		c.Delete("foo")
	}()
	go func() {
		c.Delete("bar")
	}()
	go func() {
		c.Delete("baz")
	}()

}

func TestCacheKitValues(t *testing.T) {
	c := New[string, int]()
	c.Set("hello", 1)
	c.Set("world", 2)
	c.Set("cachekit", 3)
	assert.Equal(t, 3, c.Len())

	testValues(t, c)
}

func testValues(t *testing.T, c *CacheKit[string, int]) {
	values := c.Values()
	assert.Equal(t, 3, len(values))
	assert.Contains(t, values, 1)
	assert.Contains(t, values, 2)
	assert.Contains(t, values, 3)
}

func TestCacheKitClear(t *testing.T) {
	c := New[string, int]()
	c.Set("hello", 1)
	c.Set("world", 2)
	c.Set("cachekit", 3)
	assert.Equal(t, 3, c.Len())

	c.Clear()
	assert.Equal(t, 0, c.Len())
	val, ok := c.Get("hello")
	assert.False(t, ok)
	assert.Equal(t, 0, val)
	val, ok = c.Get("world")
	assert.False(t, ok)
	assert.Equal(t, 0, val)
	val, ok = c.Get("cachekit")
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}

func TestCacheKitLen(t *testing.T) {
	c := New[string, int]()
	c.Set("hello", 1)
	c.Set("world", 2)
	c.Set("cachekit", 3)
	assert.Equal(t, 3, c.Len())

	c.Delete("hello")
	assert.Equal(t, 2, c.Len())
	c.Delete("world")
	assert.Equal(t, 1, c.Len())
	c.Delete("cachekit")
	assert.Equal(t, 0, c.Len())
}

func TestCacheKitGet(t *testing.T) {
	c := New[string, int]()
	c.Set("hello", 1)
	c.Set("world", 2)
	c.Set("cachekit", 3)
	assert.Equal(t, 3, c.Len())

	val, ok := c.Get("hello")
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	val, ok = c.Get("world")
	assert.True(t, ok)
	assert.Equal(t, 2, val)
	val, ok = c.Get("cachekit")
	assert.True(t, ok)
	assert.Equal(t, 3, val)
}

func TestCacheKitSet(t *testing.T) {
	c := New[string, int]()
	c.Set("hello", 1)
	c.Set("world", 2)
	c.Set("cachekit", 3)
	assert.Equal(t, 3, c.Len())
	v, ok := c.Get("hello")
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	v, ok = c.Get("world")
	assert.True(t, ok)
	assert.Equal(t, 2, v)
	v, ok = c.Get("cachekit")
	assert.True(t, ok)
	assert.Equal(t, 3, v)
}

func BenchmarkCacheKitSet(b *testing.B) {
	c := New[string, int]()
	for i := 0; i < b.N; i++ {
		c.Set("key"+strconv.Itoa(i), i)
	}
}

func BenchmarkCacheKitGet(b *testing.B) {
	c := New[string, int]()
	for i := 0; i < b.N; i++ {
		c.Set("key"+strconv.Itoa(i), i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = c.Get("key" + strconv.Itoa(i))
	}
}

func BenchmarkCacheKitDelete(b *testing.B) {
	c := New[string, int]()
	for i := 0; i < b.N; i++ {
		c.Set("key"+strconv.Itoa(i), i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.Delete("key" + strconv.Itoa(i))
	}
}
