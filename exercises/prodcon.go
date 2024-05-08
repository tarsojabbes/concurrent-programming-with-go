package main

import (
	"fmt"
	"math/rand"
	"time"
)

func produce(buffer chan int, done chan struct{}) {
	for i := 0; ; i++ {
		value := rand.Intn(10)
		buffer <- value
		fmt.Printf("Producer added %v to buffer\n", value)
		time.Sleep(time.Second * 1)
	}

}

func consume(buffer chan int) {
	for i := 0; ; i++ {
		value := <- buffer
		fmt.Printf("Consumer got %v from buffer\n", value)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	buffer := make(chan int, 5)
	done := make(chan struct{})

	go produce(buffer, done)
	go consume(buffer)

	<- done
}