package cache

import "time"

const defaultTTL = time.Hour

// DestructiveCache creates a cache that holds type T that can repeatedly auto-delete everything once TTL passes.
type DestructiveCache[T any] struct {
	TTL      time.Duration
	elements map[string]T
	ticker   *time.Ticker
}

func NewDestructiveCache[T any](ttl time.Duration) *DestructiveCache[T] {
	return &DestructiveCache[T]{TTL: ttl, elements: map[string]T{}}
}

func (dc *DestructiveCache[T]) Set(key string, value T) {
	dc.elements[key] = value
}

func (dc *DestructiveCache[T]) Get(key string) (T, bool) {
	value, exists := dc.elements[key]
	return value, exists
}

func (dc *DestructiveCache[T]) Delete(key string) {
	delete(dc.elements, key)
}

// Values returns all elements of type T within the cache.
func (dc *DestructiveCache[T]) Values() map[string]T {
	return dc.elements
}

// SliceValues returns a slice of elements of type T.
func (dc *DestructiveCache[T]) SliceValues() []T {
	var values []T
	for _, v := range dc.elements {
		values = append(values, v)
	}
	return values
}

// EnqueueDestruction starts repeatedly clearing all elements within the cache once TTL passes since called.
// If no TTL was given, the default of time.Hour is used. Use DequeueDestruction to stop destruction.
func (dc *DestructiveCache[T]) EnqueueDestruction() {
	if dc.TTL == 0 {
		dc.TTL = defaultTTL
	}

	dc.ticker = time.NewTicker(dc.TTL)
	for range dc.ticker.C {
		dc.elements = map[string]T{}
	}
}

// DequeueDestruction stops the repeat destroying of all elements within the cache.
func (dc *DestructiveCache[T]) DequeueDestruction() {
	dc.ticker.Stop()
	dc.ticker = nil
}
