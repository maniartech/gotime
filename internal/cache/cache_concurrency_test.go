package cache_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/maniartech/gotime/v2/internal/cache"
)

// TestConcurrentAccess is the regression guard for the "concurrent map writes"
// panic: many goroutines Set/Get/GetStrs the same small key space at once. With
// the unsynchronized map this crashes the process (and trips -race); with the
// RWMutex it is clean. Run with `go test -race ./internal/cache/`.
func TestConcurrentAccess(t *testing.T) {
	cache.Enable()

	const goroutines = 64
	const iters = 2000

	var wg sync.WaitGroup
	wg.Add(goroutines)
	for g := 0; g < goroutines; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < iters; i++ {
				key := "k" + strconv.Itoa(i%16) // shared keys => real read/write overlap
				cache.Set(key, "v"+strconv.Itoa(i%16))
				_ = cache.Get(key)
				_ = cache.GetStrs(key)
				if i%128 == 0 {
					_ = cache.IsEnabled()
				}
			}
		}()
	}
	wg.Wait()

	// The values were written many times; whatever survives must be a value we set,
	// never a torn/corrupt entry.
	if v := cache.Get("k0"); v != nil && v != "v0" {
		t.Fatalf("unexpected value for k0: %v", v)
	}
}

// TestConcurrentEnableDisable exercises the second race surface: Enable/Disable
// reassign the cache map pointer while other goroutines Get/Set. The test only
// asserts the absence of a data race / panic (values are intentionally not checked,
// since a concurrent Disable legitimately drops entries).
func TestConcurrentEnableDisable(t *testing.T) {
	cache.Enable()
	defer cache.Enable() // leave the cache enabled for any later tests

	const workers = 32
	const iters = 3000

	var wg sync.WaitGroup
	wg.Add(workers + 2)

	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for i := 0; i < iters; i++ {
				cache.Set("key"+strconv.Itoa(i%8), []string{"a", "b"})
				_ = cache.Get("key" + strconv.Itoa(i%8))
				_ = cache.GetStrs("key" + strconv.Itoa(i%8))
			}
		}()
	}
	// Two goroutines flipping the enabled state under the readers/writers.
	for f := 0; f < 2; f++ {
		go func() {
			defer wg.Done()
			for i := 0; i < iters; i++ {
				if i%2 == 0 {
					cache.Enable()
				} else {
					cache.Disable()
				}
			}
		}()
	}
	wg.Wait()
}

// BenchmarkGetParallel measures the hot read path under concurrency: after warm-up
// every call is a cached read, so this reflects how well the RWMutex read path
// scales across goroutines (the shape gotime sees when many values are formatted in
// parallel).
func BenchmarkGetParallel(b *testing.B) {
	cache.Enable()
	cache.Set("yyyy-mm-dd", []string{"2006-01-02"})
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = cache.Get("yyyy-mm-dd")
		}
	})
}

// BenchmarkSetGetMixedParallel measures a read-heavy mix (writes are rare in
// practice) under concurrency.
func BenchmarkSetGetMixedParallel(b *testing.B) {
	cache.Enable()
	keys := []string{"yyyy", "mm", "dd", "hh:ii:ss", "wwww, dd mmmm yyyy"}
	for _, k := range keys {
		cache.Set(k, []string{k})
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			k := keys[i%len(keys)]
			if i%64 == 0 { // ~1.5% writes
				cache.Set(k, []string{k})
			} else {
				_ = cache.Get(k)
			}
			i++
		}
	})
}
