package main

import (
	"fmt"
	"time"
)

type Request struct {
	id int
}

const max_capacity = 10
var req_id = 1

func exec_req(req Request) {
	fmt.Printf("Executing req id = %d\n", req.id)
	fmt.Println("Sleeping while executing")
	time.Sleep(time.Millisecond * 500)
}

func create_req() Request {
	req_id++
	return Request{id: req_id}
}

func main() {
	req_chan := make(chan Request, max_capacity)

	go func() {
		for req := range req_chan {
			exec_req(req)
		}
	}()

	for {
		if len(req_chan) < max_capacity {
			var req Request = create_req()
			req_chan <- req
		} else {
			time.Sleep(time.Second * 2)
			fmt.Println("Channel buffer reached max capacity, waiting for 1 second")
		}
	}

}