package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	wg.Add(len(channels))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//go func() {
	for _, chIn := range channels {
		go func(channelIn <-chan interface{}, ctx context.Context) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case <-channelIn:
					cancel()
				}
			}
		}(chIn, ctx)

	}
	//}()

	return chOut
}

func main() {
	//chs := make([]chan interface{}, 8)
	//
	//for i := range chs {
	//	ch := make(chan interface{})
	//	chs[i] = ch
	//}
	//
	//chsIn := make([]<-chan interface{}, len(chs))
	//for i := range chsIn {
	//	chsIn[i] = chs[i]
	//}
	//
	//chOut := or(chsIn...)
	//
	//go func() {
	//	chs[0] <- "first"
	//	chs[1] <- 2
	//	close(chs[4])
	//	chs[2] <- 4.12
	//}()
	//
	////go func() { chs[0] <- "first" }()
	////go func() { chs[1] <- 2 }()
	////go func() {
	////	chs[2] <- 4.12
	////}()
	////go func() { chs[3] <- [2]int{2, 6} }()
	////go func() { chs[4] <- "for" }()
	////go func() { chs[5] <- "A" }()
	////go func() { chs[6] <- [4]int{0, 0, 0, 0} }()
	////go func() { chs[7] <- "last" }()
	//go func() {
	//	time.Sleep(time.Second)
	//	for _, ch := range chs {
	//		_, ok := <-ch
	//		if ok {
	//			close(ch)
	//		}
	//	}
	//}()
	//
	//for val := range chOut {
	//	fmt.Printf("val: %v\n", val)
	//}

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
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone after %v", time.Since(start))
}
