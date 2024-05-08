package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	id int
	result string
}

func doTask(taskId int, task_result chan Task) {
	fmt.Printf("Executing task %v\n", taskId)
	time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
	task_result <- Task{id: taskId, result: "success"}
}

func execute(task_channel chan Task, task_result chan Task) {
	for {
		task := <- task_channel
		doTask(task.id, task_result)
	}
}

func coordinate(task_channel chan Task, task_result chan Task) {
	taskId := 0

	go func() {
		for {
			select {
			case task := <- task_result:
				fmt.Printf("Task %v terminated before timeout\n", task.id)
			case <- time.After(time.Second * 5):
				fmt.Printf("No task terminated after 5 seconds\n")
			}
		}
	}()

	for {
		task_channel <- Task{id: taskId}
		taskId++
		time.Sleep(time.Second * 1)
	}
}

func main() {
	task_channel := make(chan Task, 10)
	task_result := make(chan Task, 10)
	done := make(chan struct{})

	go coordinate(task_channel, task_channel)
	go execute(task_channel, task_result)

	<- done

}