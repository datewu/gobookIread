package memo

import (
	"context"
	"sync"
)

type entry struct {
	res   result
	ready chan struct{}
}

type result struct {
	value interface{}
	err   error
}

// Func is the type of the function can be memorize.
type Func func(string) (interface{}, error)

// Memo caches the results of calling a Func
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

// New wraper the Func to Memo
func New(ctx context.Context, f Func) *Memo {
	m := &Memo{f: f, cache: make(map[string]*entry)}
	go func() {
		<-ctx.Done()
		m.clear()
	}()
	return m
}

// clear all cache entry
func (m *Memo) clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cache = make(map[string]*entry)
}

// Get the result
func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	e := m.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		m.cache[key] = e
		m.mu.Unlock()

		e.res.value, e.res.err = m.f(key)
		close(e.ready) // broadcase ready conditon.
	} else {
		m.mu.Unlock()
		<-e.ready // block on the nil channel, wait for it close
	}
	return e.res.value, e.res.err
}
