package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Resource struct {
	id int
	available chan bool
}

func (r *Resource) get(tid int) {
	r.available <- false
	fmt.Printf("Thread %v got resource %v\n", tid, r.id)
}

func (r *Resource) release(tid int) {
	<- r.available
	fmt.Printf("Thread %v relesead resource %v\n", tid, r.id)
}

func consumeResource(resources []Resource) {
	for i := 0; ; i++ {
		go func(i int) {id := rand.Intn(len(resources))
		resources[id].get(i)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		resources[id].release(i)}(i)
	}
}

func main(){
	r0 := Resource{id: 0, available: make(chan bool, 1)}
	r1 := Resource{id: 1, available: make(chan bool, 1)}
	r2 := Resource{id: 2, available: make(chan bool, 1)}
	r3 := Resource{id: 3, available: make(chan bool, 1)}
	r4 := Resource{id: 4, available: make(chan bool, 1)}

	resources := []Resource {r0, r1, r2, r3, r4}

	consumeResource(resources)

}

