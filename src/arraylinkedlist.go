package main

import (
	"errors"
)

type ANode struct {
	next  int
	value int
}

type ALinkedList struct {
	fempty int
	fvalue int
	list   []ANode
}

func NewALinkedList() *ALinkedList {
	list := make([]ANode, 5, 5)
	for i, _ := range list[:] {
		if i < (len(list) - 1) {
			list[i] = ANode{
				next:  i + 1,
				value: 0,
			}
		} else {
			list[i] = ANode{
				next:  -1,
				value: 0,
			}
		}
	}

	return &ALinkedList{
		list:   list[:],
		fempty: 0,
		fvalue: -1,
	}
}

func (al *ALinkedList) resize() error {
	al.list = append(al.list, ANode{
		next:  -1,
		value: 0,
	})
	al.fempty = al.size() - 1
	return nil
}

func (al *ALinkedList) Push(x int) error {
	flag := 0
	if al.fvalue < 0 {
		al.fvalue = al.fempty
		al.fempty = al.list[al.fempty].next
		al.list[al.fvalue] = ANode{
			value: x,
			next:  -1,
		}
	} else {
		if al.fempty == -1 {
			al.resize()
		}
		flag = al.fvalue
		al.list[al.fempty].value = x
		al.fvalue, al.fempty = al.fempty, al.list[al.fempty].next
		al.list[al.fvalue].next = flag
	}
	return nil
}

func (al *ALinkedList) Pop() (int, error) {
	flag := 0
	if al.fvalue < 0 {
		return 0, errors.New("empty")
	} else {
		flag = al.fempty
		al.fempty, al.fvalue = al.fvalue, al.list[al.fvalue].next
		al.list[al.fempty].next = flag
		return al.list[al.fempty].value, nil
	}
}

func (al *ALinkedList) size() int {
	return len(al.list)
}

func (al *ALinkedList) isEmpty() bool {
	return al.fvalue >= 0
}

func (al *ALinkedList) removeALl() {
	al.fempty = 0
	al.fvalue = -1
}

func main() {
	al := NewALinkedList()
	al.resize()
	println(al)
}
