package errx

import "sync"

var (
	registry = map[string]struct{}{}
	mu       sync.Mutex
)

func register(code string) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := registry[code]; exists {
		panic("errorx: duplicate error code " + code)
	}

	registry[code] = struct{}{}
}
