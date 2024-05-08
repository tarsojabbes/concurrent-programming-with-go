package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Dish struct {
	dish_type string
}

func do_dish(kitchen_counter chan Dish) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	kitchen_counter <- Dish{}
}

func cooker(kitchen_counter chan Dish) {
	for {
		if len(kitchen_counter) < cap(kitchen_counter) {
			go do_dish(kitchen_counter)
		} else {
			fmt.Println("Kitchen counter is full")
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		}
		
	}
}

func consumer(kitchen_counter chan Dish, mutex chan struct{}) {
	for {
		mutex <- struct{}{}
		<- kitchen_counter
		<- mutex
		fmt.Println("Consuming dish")
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

func main() {
	kitchen_counter := make(chan Dish, 10)
	mutex := make(chan struct{}, 1)
	done := make(chan struct{})

	go cooker(kitchen_counter)

	go func(){
		for i := 0; i < 10; i++{
			go consumer(kitchen_counter, mutex)
		}
	}()

	<- done
}
