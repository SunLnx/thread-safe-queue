package queue

import (
	"fmt"
	`math/rand`
	"testing"
)

func TestQueue(t *testing.T) {
	var goroutine int = 10
	var q *Queue = NewQueue()
	var ch chan byte = make(chan byte, goroutine)
	for i := 0; i < goroutine; i++ {
		go func(value interface{}, ch chan byte) {
			q.Enqueue(rand.Intn(100))
			ch <- '0'

		}(i, ch)
	}

	var sum int = 0
	for {
		select {
		case <-ch:
			sum++
			if sum == goroutine {
				//q.SnapShot()
				for value := q.Dequeue(); value != nil; value = q.Dequeue() {
					//fmt.Println(value)
				}
				fmt.Println(q.Len())
				//q.SnapShot()
				return
			}
		}
	}

}
