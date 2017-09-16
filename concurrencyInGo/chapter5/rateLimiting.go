package main

import (
	"context"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func limitDemo() {
	defer log.Println("Done")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	api := open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := api.readFile(context.Background())
			if err != nil {
				log.Println("cannot ReadFile:", err)
			}
			log.Println("ReadFile")
		}()

	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := api.resolveAddress(context.Background())
			if err != nil {
				log.Println("cannot ResolveAddress:", err)
			}
			log.Println("resolveAddress")
		}()

	}
	wg.Wait()
}

type rateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

type multiLimiter struct {
	limiters []rateLimiter
}

func (l *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (l *multiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
}

func newMultiLimiter(limiters ...rateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit)
	return &multiLimiter{limiters}
}

type apiConnection struct {
	networkLimit, diskLimit, apiLimit rateLimiter
}

func per(evertCont int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(evertCont))
}

func open() *apiConnection {
	apiLimit := newMultiLimiter(
		rate.NewLimiter(per(2, time.Second), 2),
		rate.NewLimiter(per(10, time.Minute), 10),
	)

	diskLimit := newMultiLimiter(
		rate.NewLimiter(rate.Limit(1), 1),
	)

	networkLimit := newMultiLimiter(
		rate.NewLimiter(per(3, time.Second), 5),
	)
	return &apiConnection{
		networkLimit,
		diskLimit,
		apiLimit,
	}
}

func (a *apiConnection) readFile(ctx context.Context) error {
	err := newMultiLimiter(a.apiLimit, a.diskLimit).Wait(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (a *apiConnection) resolveAddress(ctx context.Context) error {
	err := newMultiLimiter(a.apiLimit, a.networkLimit).Wait(ctx)
	if err != nil {
		return err
	}
	return nil
}
