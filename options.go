package temporal

type options struct {
	cache map[string]string
}

// Options is the global options for the temporal package. You can use this to
// disable the cache and other options.
//
// For example:
//
//	temporal.Options.DisableCache()
//	temporal.Options.EnableCache()
var Options *options

func init() {
	Options = &options{
		cache: map[string]string{},
	}
}

// DisableCache disables the cache. This is useful for testing or
// when you want to make sure that the cache is not used. Generally,
// many applications do not need to disable the cache due to the limited
// number of datetime formats that are used. However, if you are using
// unknown number of datetime formats, then you may want to disable
// the cache. The cache is enabled by default. To reenable the cache,
// use the EnableCache function.
func (o *options) DisableCache() *options {
	o.cache = nil
	return o
}

// EnableCache enables the cache. By default, the cache is enabled.
// This function useful when you want to enable the cache after it has been
// disabled. See the DisableCache function for more information.
func (o *options) EnableCache() *options {
	o.cache = map[string]string{}
	return o
}

// IsCacheEnabled returns true if the cache is enabled. The cache is
// enabled by default. See the DisableCache function for more information.
func (o *options) IsCacheEnabled() bool {
	return o.cache != nil
}
