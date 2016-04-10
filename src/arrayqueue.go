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
	// q := &ArrayQueueA{}
	// return q //直接return
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
	// if lq.checkEnqueue() == false { //!
	// 	return errors.New("queue is full")
	// } else { // 移出
	// 	lq.list[(lq.head % 5)] = v
	// 	lq.head = lq.head + 1
	// 	return nil
	// }
}

func (lq *ArrayQueueA) Dequeue() (int, error) {
	if !lq.checkDequeue() {
		return 0, errors.New("queue is empty")
	}
	v := lq.list[(lq.tail % 5)] //
	lq.tail = lq.tail + 1
	return v, nil
	// if lq.checkDequeue() == false {
	// 	return 0, errors.New("queue is empty")
	// } else {
	// 	v := lq.list[(lq.tail % 5)] //
	// 	lq.tail = lq.tail + 1
	// 	return v, nil
	// }
}

func (lq *ArrayQueueA) checkDequeue() bool {
	return lq.head > lq.tail
	// if lq.head > lq.tail { // return > bool
	// 	return true
	// } else {
	// 	return false
	// }
}

func (lq *ArrayQueueA) checkEnqueue() bool {
	return !((lq.head%5) == (lq.tail%5) && (lq.head > lq.tail))
	// if (lq.head%5) == (lq.tail%5) && (lq.head > lq.tail) {
	// 	return false
	// } else {
	// 	return true
	// }
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
