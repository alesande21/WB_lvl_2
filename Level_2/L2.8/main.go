package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	var once sync.Once
	go func() {
		//defer close(out)
		var wg sync.WaitGroup
		wg.Add(len(channels))
		for _, ch := range channels {
			go func(in <-chan interface{}) {
				defer wg.Done()
				select {
				case <-in:
					once.Do(func() { close(out) })
				}
			}(ch)
		}
		go func() {
			defer close(out)
			wg.Wait()
		}()
	}()

	return out
}

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(4*time.Second),
		sig(2*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone after %v", time.Since(start))
}
