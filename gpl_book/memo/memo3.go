package memo

import (
	"io/ioutil"
	"net/http"
	"sync"
)

// Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New wraper Func to Memo
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get memory
func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	res, ok := m.cache[key]
	m.mu.Unlock()
	if !ok {
		res.value, res.err = m.f(key)

		// Between the two critical sections, several goroutines
		// may race to compute f(key) and update the map
		m.mu.Lock()
		m.cache[key] = res
		m.mu.Unlock()
	}
	return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
