package memo

import "sync"

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type result struct {
	value interface{}
	err   error
}

// A Memo caches the results of calling a Func
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// New lol
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Get NOTE: not concurrency-safe!
func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	e := m.cache[key]
	if e == nil {
		// This the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready conditon.
		e = &entry{ready: make(chan struct{})}
		m.cache[key] = e
		m.mu.Unlock()
		e.res.value, e.res.err = m.f(key)
		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key
		m.mu.Unlock()
		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}
