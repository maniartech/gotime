package dateutils

// This file contains constants that maps django date format to go date format

var cache = map[string]string{}

// DisableCache disables the cache. This is useful for testing or
// when you want to make sure that the cache is not used. Generally,
// many applications do not need to disable the cache due to the limited
// number of datetime formats that are used. However, if you are using
// unknown number of datetime formats, then you may want to disable
// the cache. The cache is enabled by default. To reenable the cache,
// use the EnableCache function.
func DisableCache() {
	cache = nil
}

// EnableCache enables the cache. By default, the cache is enabled.
// This function useful when you want to enable the cache after it has been
// disabled. See the DisableCache function for more information.
func EnableCache() {
	cache = map[string]string{}
}
