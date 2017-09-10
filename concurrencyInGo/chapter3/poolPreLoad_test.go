package main

// ➜  chapter3 git:(master) ✗ go test  -benchtime=10s -bench=BenchmarkNetworkRequestPool
// goos: darwin
// goarch: amd64
// BenchmarkNetworkRequestPool-4               5000           3136301 ns/op
// PASS
// ok      _/Users/rack/golangCodeBase/books/concurrencyInGo/chapter3      53.824s
// ➜  chapter3 git:(master) ✗
// ➜  chapter3 git:(master) ✗
// ➜  chapter3 git:(master) ✗ go test  -benchtime=10s -bench=BenchmarkNetworkRequestWithoutPool
// goos: darwin
// goarch: amd64
// BenchmarkNetworkRequestWithoutPool-4          10        1002297174 ns/op
// PASS
// ok      _/Users/rack/golangCodeBase/books/concurrencyInGo/chapter3      31.111s

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}

}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			connectToService()
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()
	return &wg
}

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()

	warmDaemonStarted := warmStartNetworkDaemon()
	warmDaemonStarted.Wait()
}

func BenchmarkNetworkRequestWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannnot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}

}

func BenchmarkNetworkRequestPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:9090")
		if err != nil {
			b.Fatalf("cannnot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}

}

func warmSerivceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 20; i++ {
		p.Put(p.New())

	}
	return p
}

func warmStartNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmSerivceConnCache()
		server, err := net.Listen("tcp", "localhost:9090")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			//			connectToService()
			serviceConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(serviceConn)
			conn.Close()
		}
	}()
	return &wg
}
