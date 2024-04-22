package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const n = 10

func draw_and_sleep(tid int, wg *sync.WaitGroup) {
	defer wg.Done()
	sleep_time := rand.Int() * 5 // probably not how we want to draw a random number between 0 and 5
	fmt.Printf("Thread %d will be sleeping for %d seconds\n", tid, sleep_time)
	time.Sleep(time.Second * time.Duration(sleep_time))
	fmt.Printf("Thread %d done sleeping\n", tid)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go draw_and_sleep(i, &wg)
	}

	wg.Wait()
	fmt.Printf("n is %d\n", n)
}