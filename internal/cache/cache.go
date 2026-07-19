package cache

import "sync"

// mu guards both the cache map and its enabled/disabled state.
//
// It is an RWMutex because Get and GetStrs sit on the hot path — they are called
// once per datetime-layout conversion — and gotime is routinely used from
// concurrent code (formatting or parsing many values in parallel). The cache is a
// plain Go map, and an unsynchronized map panics the whole process with
// "fatal error: concurrent map writes" (or a concurrent read/write) the moment two
// goroutines touch it at once. RWMutex lets any number of readers run concurrently
// while writes are exclusive.
//
// Writes are rare: the set of datetime layouts a program uses is small and
// stabilizes almost immediately, after which every call is a contention-light
// read. Enable/Disable also take the write lock, which additionally makes the
// map-pointer reassignment they perform safe against a concurrent Get/Set (a
// second race the previous unsynchronized version had).
var mu sync.RWMutex

var cache map[string]interface{}

func init() {
	cache = map[string]interface{}{}
}

// Disable disables the cache. This is useful for testing or
// when you want to make sure that the cache is not used. Generally,
// many applications do not need to disable the cache due to the limited
// number of datetime formats that are used. However, if you are using
// unknown number of datetime formats, then you may want to disable
// the cache. The cache is enabled by default. To reenable the cache,
// use the Enable function.
func Disable() {
	mu.Lock()
	cache = nil
	mu.Unlock()
}

// Enable enables the cache. By default, the cache is enabled.
// This function is useful when you want to enable the cache after it has been
// disabled. See the Disable function for more information.
func Enable() {
	mu.Lock()
	cache = map[string]interface{}{}
	mu.Unlock()
}

// IsEnabled returns true if the cache is enabled. The cache is
// enabled by default. See the Disable function for more information.
func IsEnabled() bool {
	mu.RLock()
	defer mu.RUnlock()
	return cache != nil
}

// Set sets the cache value for the given key. If the cache is disabled,
// this function does nothing.
func Set(key string, value interface{}) {
	mu.Lock()
	if cache != nil {
		cache[key] = value
	}
	mu.Unlock()
}

// Get returns the cache value for the given key. If the cache is disabled,
// this function returns nil.
func Get(key string) interface{} {
	mu.RLock()
	defer mu.RUnlock()
	if cache == nil {
		return nil
	}
	return cache[key]
}

func GetStrs(key string) []string {
	mu.RLock()
	defer mu.RUnlock()
	if cache == nil {
		return nil
	}
	value, exists := cache[key]
	if !exists {
		return nil
	}
	switch v := value.(type) {
	case []string:
		return v
	case string:
		return []string{v}
	default:
		return nil
	}
}
