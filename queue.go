package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	capacity    int
	queuedItems chan string
}

func (queue *Queue) Insert(item string) error {
	if len(queue.queuedItems) < queue.capacity {
		queue.queuedItems <- item

		return nil
	}

	return errors.New("Queue is full")
}

func (queue *Queue) RemoveBatch() *[]string {
	fmt.Println("RemoveBatch")

	count := len(queue.queuedItems)
	items := make([]string, count)

	fmt.Printf("Count: %d", count)
	if count > 0 {
		for i := 0; i < count; i++ {
			item := <-queue.queuedItems
			items[i] = item
		}
	}

	return &items
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		capacity:    capacity,
		queuedItems: make(chan string, capacity),
	}
}
