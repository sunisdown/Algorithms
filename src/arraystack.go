package main

import (
	"errors"
)

type ArrayStack struct {
	list []int
}

func NewArrayStack() *ArrayStack {
	q := &ArrayStack{
		list: make([]int, 0, 5),
	}
	return q
}

func (as *ArrayStack) size() int {
	return len(as.list)
}

// 出栈
func (as *ArrayStack) Pop() (int, error) {
	if as.isEmpty() {
		return 0, errors.New("stack is empty")
	} else {
		x := as.list[as.size()-1]
		as.list = as.list[:(as.size() - 1)]
		return x, nil
	}
}

// 入栈
func (as *ArrayStack) Push(x int) error {
	if as.isFull() {
		return errors.New("stack is full")
	} else {
		as.list = as.list[:(as.size() + 1)]
		println("llll\n", as.size())
		as.list[as.size()-1] = x
		return nil
	}
}

// 判断是否为空
func (as *ArrayStack) isEmpty() bool {
	return as.size() == 0
}

// 判断栈是否为空
func (as *ArrayStack) isFull() bool {
	return as.size() == cap(as.list)
}
