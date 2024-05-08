package main

import (
	"fmt"
	"math/rand"
	"time"
)

type File struct {
	name string
	size int
}

type Request struct {
	filename string
}

func server(transfer_file_request chan Request, file_buffer chan File) {
	for {
		request := <- transfer_file_request
		fmt.Printf("[SERVER] -- Request received from client looking for file %v\n", request.filename)
		time.Sleep(time.Duration(rand.Intn(6)))
		file_buffer <- File{name: request.filename, size: rand.Intn(10) * 169}
		time.Sleep(time.Second * 5)
	}
}

func client(ip string, transfer_file_request chan Request, file_buffer chan File) {
	for {
		transfer_file_request <- Request{filename: "abc.txt"}
		fmt.Printf("[CLIENT - %v] -- Request sent\n", ip)
		file := <- file_buffer
		fmt.Printf("[CLIENT - %v] -- File %v received by client | Total bytes: %v\n", ip, file.name, file.size)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	transfer_file_request := make(chan Request, 10)
	file_buffer := make(chan File, 10)
	done := make(chan struct{})

	go client("172.92.10.30", transfer_file_request, file_buffer)
	go client("192.112.30.15", transfer_file_request, file_buffer)
	go client("10.10.63.244", transfer_file_request, file_buffer)
	go server(transfer_file_request, file_buffer)

	<- done

}