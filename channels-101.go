package main

import "fmt"

func main() {

    messages := make(chan string) // using make to create a new string channel

    go func() { messages <- "ping" }() // closure for putting "ping" to channel

    msg := <-messages // reading channel
    fmt.Println(msg)
}