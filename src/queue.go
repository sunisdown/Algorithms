package main

import "sync"
import "fmt"

type Node struct {
	item interface{}
	next *Node
}

type Queue struct {
	head  *Node
	tail  *Node
	count int
	lock  *sync.Mutex
}

func NewQueue() *Queue {
	q := &Queue{}
	q.lock = &sync.Mutex{}
	return q
}

func (self *Queue) size() int {
	self.lock.Lock()
	defer self.lock.Unlock()
	return self.count
}

func (self *Queue) isEmpty() bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	return self.count == 0
}

func (self *Queue) Enqueue(item interface{}) {
	self.lock.Lock()
	defer self.lock.Unlock()

	node := &Node{item: item}

	if self.tail == nil {
		self.tail = node
		self.head = node
	} else {
		self.tail.next = node
		self.tail = node
	}
	self.count++
}

func (self *Queue) Dequeue() *Node {
	self.lock.Lock()
	defer self.lock.Unlock()
	item := self.head
	self.head = self.head.next
	if self.head == nil {
		self.tail = nil
	}
	self.count--
	return item
}

func main() {
	var queue *Queue = NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.head.item)
	queue.Dequeue()
	fmt.Println(queue.head.item)
}
