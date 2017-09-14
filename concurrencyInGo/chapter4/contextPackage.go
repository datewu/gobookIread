package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type ctxKey int

const (
	lo        = "en/us"
	a  ctxKey = iota
	b
	c
)

func main() {
	//withoutContext()
	//usingContext()

	ctx := context.WithValue(context.Background(), a, "name")
	ctx = context.WithValue(ctx, b, "email")
	ctx = context.WithValue(ctx, c, "address")

	fmt.Println(ctx.Value(a), ctx.Value(1), ctx.Value(b), ctx.Value(c))

}

func usingContext() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := greeting(ctx); err != nil {
			fmt.Println("cannot print greeting", err)
			cancel()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := farewell(ctx); err != nil {
			fmt.Println("cannot print farewell", err)
		}

	}()
	wg.Wait()

}

func greeting(ctx context.Context) error {
	gr, err := genGr(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!", gr)
	return nil
}

func farewell(ctx context.Context) error {
	fr, err := genFr(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!", fr)
	return nil
}

func genGr(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	switch locale, err := loc(ctx); {
	case err != nil:
		return "", err
	case locale == lo:
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFr(ctx context.Context) (string, error) {
	switch locale, err := loc(ctx); {
	case err != nil:
		return "", err
	case locale == lo:
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func loc(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(time.Minute):
	}
	return lo, nil

}

func withoutContext() {
	var wg sync.WaitGroup
	done := make(chan interface{})
	defer close(done)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(done); err != nil {
			fmt.Println(err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(done); err != nil {
			fmt.Println(err)
			return
		}
	}()
	wg.Wait()

}

func printGreeting(done <-chan interface{}) (err error) {
	if greeting, err := genGreeting(done); err == nil {
		fmt.Printf("%s world!", greeting)
	}
	return

}
func printFarewell(done <-chan interface{}) (err error) {
	if farewell, err := genFarewell(done); err == nil {
		fmt.Printf("%s world!", farewell)
	}
	return
}

func genGreeting(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == lo:
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == lo:
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}
func locale(done <-chan interface{}) (string, error) {
	select {
	case <-done:
		return "", fmt.Errorf("canceled")
	case <-time.After(1 * time.Minute):
	}
	return lo, nil
}
