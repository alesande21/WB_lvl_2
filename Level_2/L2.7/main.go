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
		for _, ch := range channels {
			go func(in <-chan interface{}) {
				select {
				case <-in:
					once.Do(func() { close(out) })
				}
			}(ch)
		}
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
	<-or(sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second))
	fmt.Printf("fone after %v", time.Since(start))

}
