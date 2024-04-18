package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	<- done // this line will block the main goroutine from finishing before worker goroutine is done
}