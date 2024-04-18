package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }() // After 1 second writes "one" on c1

    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }() // After 2 seconds writes "two" on c2

	// select can be combined with channels and go routines
	// the first one to arrive will be selectec.
	// we are using a for loop to demonstrate that, on the first loop
	// c1 receives "one", and on the second loop c2 receives "two"
    for i := 0; i < 2; i++ {
        select { 
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}