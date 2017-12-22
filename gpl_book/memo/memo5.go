package memo

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

type request struct {
	key      string
	response chan<- result
}

// Memo caches the results of calling a Func
type Memo struct {
	requests chan request
}

// New wraper the Func to Memo
func New(f Func) *Memo {
	m := &Memo{make(chan request)}
	go m.server(f)
	return m
}

// Get the cache
func (m *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	m.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

// Close memo Server
func (m *Memo) Close() {
	close(m.requests)
}

func (m *Memo) server(f Func) {
	cache := make(map[string]*entry)

	for req := range m.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
