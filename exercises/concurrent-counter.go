package main

import (
	"fmt"
	"sync"
)

var counter = 0

func increment(mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()

}

func decrement(mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done() 
	mutex.Lock()
	counter--
	mutex.Unlock()
}
	

func getValue(mutex *sync.Mutex) int {
	defer mutex.Unlock() 
	mutex.Lock()
	return counter
}


func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go increment(&mutex, &wg)
	go increment(&mutex, &wg)

	wg.Wait()

	fmt.Println("Valor final do contador:", getValue(&mutex))

}