package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type startGoroutineFn func(
	done <-chan interface{},
	pulseInterval time.Duration,
) (heartbeat <-chan interface{})

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)
	// simpleSteWard()
	fnSteWard()
}

func newSteward(timeout time.Duration,
	startGoroutine startGoroutineFn) startGoroutineFn {

	return func(
		done <-chan interface{},
		pulseInterval time.Duration,

	) <-chan interface{} {
		heartbeat := make(chan interface{})
		go func() {
			defer close(heartbeat)

			var wardDone chan interface{}
			var wardHeartbeat <-chan interface{}

			startWard := func() {
				wardDone = make(chan interface{})
				wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2)
			}

			startWard()
			pulse := time.Tick(pulseInterval)

		monitorLoop:
			for {
				timeoutSignal := time.After(timeout)
				for {
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat:
						continue monitorLoop
					case <-timeoutSignal:
						log.Println("steward: ward unhealth; restarting")
						close(wardDone)
						startWard()
						continue monitorLoop
					case <-done:
						return
					}
				}
			}

		}()

		return heartbeat
	}
}

func simpleSteWard() {
	doWork := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
		log.Println("ward: Hello, i'm irresponsible!")
		go func() {
			<-done
			log.Println("ward: I am halting")
		}()
		return nil

	}

	doWorkWithSteward := newSteward(5*time.Second, doWork)
	done := make(chan interface{})

	time.AfterFunc(14*time.Second, func() {
		log.Println("main: halting steward and ward.")
		close(done)
	})

	for v := range doWorkWithSteward(done, 3*time.Second) {

		log.Println(v)

	}
	log.Println("Done")
}

func take(
	done <-chan interface{},
	valueStream <-chan interface{},
	num int,
) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func or(c, done <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()

	return valStream
}

func bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)
		for {

			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if ok == false {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}
			for val := range or(stream, done) {
				select {
				case valStream <- val:
				case <-done:
				}

			}
		}
	}()
	return valStream
}

func fnSteWard() {
	doWorkFn := func(
		done <-chan interface{},
		intList ...int,
	) (startGoroutineFn, <-chan interface{}) {

		intChanStream := make(chan (<-chan interface{}))
		intStream := bridge(done, intChanStream)

		doWork := func(
			done <-chan interface{},
			pulseInterval time.Duration,
		) <-chan interface{} {
			iStream := make(chan interface{})
			heartbeat := make(chan interface{})

			go func() {
				defer close(iStream)

				select {
				case intChanStream <- iStream:
				case <-done:
					return
				}

				pulse := time.Tick(pulseInterval)
				for {
				valueLoop:
					for _, intVal := range intList {
						if intVal < 0 {
							log.Println("negative value:", intVal)
							return
						}
						for {
							select {
							case <-pulse:
								select {
								case heartbeat <- struct{}{}:
								default:
								}
							case iStream <- intVal:
								continue valueLoop
							case <-done:
								return
							}

						}

					}

				}

			}()
			return heartbeat

		}

		return doWork, intStream
	}

	done := make(chan interface{})
	defer close(done)

	doWork, outputs := doWorkFn(done, 1, 3, 5, -5, 7, 8)
	doWithSteward := newSteward(time.Millisecond, doWork)
	doWithSteward(done, time.Hour)

	for i := range take(done, outputs, 7) {
		fmt.Println("Received:", i)
	}

}
