package main

import "fmt"

func main() {

    messages := make(chan string, 2)


	/* Because this channel is buffered, we can send these values 
	into the channel without a corresponding concurrent receive. */
    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}