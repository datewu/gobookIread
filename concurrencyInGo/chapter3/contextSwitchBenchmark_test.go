/*
➜  chapter3 git:(master) ✗ go test -bench=. -cpu=1
goos: darwin
goarch: amd64
BenchmarkContextSwitch  10000000               182 ns/op
PASS
ok      _/Users/rack/golangCodeBase/books/concurrencyInGo/chapter3      2.013s
➜  chapter3 git:(master) ✗ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkContextSwitch-4        10000000               238 ns/op
PASS
ok      _/Users/rack/golangCodeBase/books/concurrencyInGo/chapter3      2.633s
*/
package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}

	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token

		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c

		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
