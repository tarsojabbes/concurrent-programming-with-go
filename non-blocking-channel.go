package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

	// Immediately take the default case without blocking
	// once there's nothing to read on messages
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }


	// Immediately take the default case without blocking
	// because messages is not a bufferized channel
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

	// Here we attempt non-blocking receives on both messages and signals.
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}