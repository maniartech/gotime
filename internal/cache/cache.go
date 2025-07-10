package cache

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
	cache = nil
}

// Enable enables the cache. By default, the cache is enabled.
// This function is useful when you want to enable the cache after it has been
// disabled. See the Disable function for more information.
func Enable() {
	cache = map[string]interface{}{}
}

// IsEnabled returns true if the cache is enabled. The cache is
// enabled by default. See the Disable function for more information.
func IsEnabled() bool {
	return cache != nil
}

// Set sets the cache value for the given key. If the cache is disabled,
// this function does nothing.
func Set(key string, value interface{}) {
	if cache == nil {
		return
	}
	cache[key] = value
}

// Get returns the cache value for the given key. If the cache is disabled,
// this function returns nil.
func Get(key string) interface{} {
	if cache == nil {
		return nil
	}
	return cache[key]
}

func GetStrs(key string) []string {
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
