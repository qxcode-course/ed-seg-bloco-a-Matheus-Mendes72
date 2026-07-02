package main

import (
	"fmt"
)

func main() {
	queue := NewQueue[string]()
	for i := 'A'; i <= 'P'; i++ {
		queue.Enqueue(string(i))
	}

	for i := 0; i < 15; i++ {
		var MM, NN int
		fmt.Scan(&MM, &NN)

		time1 := queue.Dequeue()
		time2 := queue.Dequeue()

		if MM > NN {
			queue.Enqueue(time1)
		} else {
			queue.Enqueue(time2)
		}
	}

	fmt.Println(queue.Dequeue())
}
