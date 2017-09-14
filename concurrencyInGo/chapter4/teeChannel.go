package main

import "fmt"

func teeDemo() {
	tee := func(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})

		go func() {
			defer close(out1)
			defer close(out2)

			for val := range orDone(done, in) {
				var a, b = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case <-done:
					case a <- val:
						a = nil
					case b <- val:
						b = nil
					}

				}

			}
		}()

		return out1, out2
	}

	done := make(chan interface{})
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})

		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()

		return valueStream
	}
	demo := take(done, repeat(done, 4, 5), 20)
	a, b := tee(done, demo)

	go func() { // NOTE: must in a different goroutine

		for v := range a {
			fmt.Println(v, "a")

		}
	}()
	fmt.Println("=========")
	for v := range b {
		fmt.Println(v, "b")

	}
}
