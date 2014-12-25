package main

import (
	"fmt"
	"sync"
)


type item struct {
	priority int
	value    interface{}
}

type PQueue struct {
	sync.RWMutex
	items      []*item
	elemsCount int
}

func newItem(priority int, value interface{}) *item {
	return &item{
		priority: priority,
		value:    value,
	}
}

func (i *item) String() string {
	return fmt.Sprintf("<item value:%s priority:%d>", i.value, i.priority)
}

func NewPQueue() *PQueue {

	items := make([]*item, 1)
	items[0] = nil // Heap queue first element should always be nil

	return &PQueue{
		items:      items,
		elemsCount: 0,
	}
}

func (pq *PQueue) insert(priority int, value interface{}) {
	item := newItem(priority, value)

	pq.Lock()
	pq.items = append(pq.items, item)
	pq.elemsCount += 1
	pq.swim(pq.size())
	pq.Unlock()
}

func (pq *PQueue) deletemin() (interface{}, int) {
	pq.Lock()
	defer pq.Unlock()

	if pq.size() < 1 {
		return nil, 0
	}

	var max *item = pq.items[1]

	pq.exch(1, pq.size())
	pq.items = pq.items[0:pq.size()]
	pq.elemsCount -= 1
	pq.sink(1)

	return max.value, max.priority
}

func (pq *PQueue) min() (interface{}, int) {
	pq.RLock()
	defer pq.RUnlock()

	if pq.size() < 1 {
		return nil, 0
	}

	headValue := pq.items[1].value
	headPriority := pq.items[1].priority

	return headValue, headPriority
}

func (pq *PQueue) Size() int {
	pq.RLock()
	defer pq.RUnlock()
	return pq.size()
}

func (pq *PQueue) size() int {
	return pq.elemsCount
}

func (pq *PQueue) less(i, j int) bool {
	return pq.items[i].priority > pq.items[j].priority
}

func (pq *PQueue) isEmpty() bool {
	return pq.elemsCount == 0
}

func (pq *PQueue) exch(i, j int) {
	var tmpItem *item = pq.items[i]

	pq.items[i] = pq.items[j]
	pq.items[j] = tmpItem
}

func (pq *PQueue) swim(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.exch(k/2, k)
		k = k / 2
	}

}

func (pq *PQueue) sink(k int) {
	for 2*k <= pq.size() {
		var j int = 2 * k

		if j < pq.size() && pq.less(j, j+1) {
			j++
		}

		if !pq.less(k, j) {
			break
		}

		pq.exch(k, j)
		k = j
	}
}


func main() {
	pq := NewPQueue()
	pq.insert(2, "world")
	pq.insert(1, "hello")
	pq.insert(4, "你好")
	head_value, head_pri := pq.min()
	fmt.Println(head_value, head_pri)
	head_value, head_pri = pq.deletemin()
	fmt.Println(head_value, head_pri)
}
