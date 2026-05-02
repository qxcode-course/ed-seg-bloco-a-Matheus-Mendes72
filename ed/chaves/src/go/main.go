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
		var gols1, gols2 int
		fmt.Scan(&gols1, &gols2)

		time1 := queue.Dequeue()
		time2 := queue.Dequeue()

		if gols1 > gols2 {
			queue.Enqueue(time1)
		} else {
			queue.Enqueue(time2)
		}
	}

	fmt.Println(queue.Dequeue())
}
