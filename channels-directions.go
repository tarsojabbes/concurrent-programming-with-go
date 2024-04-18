package main

import "fmt"

func ping(pings chan <- string, msg string) { // chan <- type defines a write-only channel
	pings <- msg
}

func pong(pings <- chan string, pongs chan <- string) { // <- chan type defines a read-only channel
	msg := <- pings
	pongs <- msg
}

func main() {
	pings := make(chan string,1) // including 1 as a second parameter defines channel capacity
	pongs := make(chan string,1)

	ping(pings, "sent message")
	pong(pings, pongs)
	fmt.Println((<- pongs))
}