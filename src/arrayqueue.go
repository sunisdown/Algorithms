package main

import (
	"errors"
)

type ArrayQueueA struct {
	head int
	tail int
	list [5]int
}

func NewArrayQueueA() *ArrayQueueA {
	return &ArrayQueueA{}
}

func (lq *ArrayQueueA) size() int { // lq
	return lq.head - lq.tail
}

func (lq *ArrayQueueA) Enqueue(v int) error {
	if !lq.checkEnqueue() {
		return errors.New("queue is full")
	}
	lq.list[(lq.head % 5)] = v
	lq.head = lq.head + 1
	return nil
}

func (lq *ArrayQueueA) Dequeue() (int, error) {
	if !lq.checkDequeue() {
		return 0, errors.New("queue is empty")
	}
	v := lq.list[(lq.tail % 5)] //
	lq.tail = lq.tail + 1
	return v, nil
}

func (lq *ArrayQueueA) checkDequeue() bool {
	return lq.head > lq.tail
}

func (lq *ArrayQueueA) checkEnqueue() bool {
	return !((lq.head%5) == (lq.tail%5) && (lq.head > lq.tail))
}

func main() {
	q := NewArrayQueueA()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	ev, err := q.Dequeue()
	println(ev)
	ev, err = q.Dequeue()
	println(ev, err)
}
