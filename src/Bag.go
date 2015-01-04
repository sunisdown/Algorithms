package main

import "fmt"

type Node struct {
	item interface{}
	next *Node
}

type Bag struct {
	size  int
	first *Node
}

func newBag() *Bag {
	return &Bag{
		size:  0,
		first: nil,
	}
}

func (self *Bag) isEmpty() bool {
	return self.first == nil
}

func (self *Bag) add(item interface{}) {
	if self.first == nil {
		self.first = &Node{
			item: item,
		}
		self.size = 1
	} else {
		old_first := self.first
		self.first = &Node{
			item: item,
		}
		self.first.next = old_first
		self.size++
	}
}

func main() {
	bag := newBag()
	fmt.Println(bag)
	bag.add("hello")
	fmt.Println(bag.first.item)
}
