package main

import (
	"fmt"
	"sync"
)

const (
	Red =  true
	Black = false
)

type Node struct {
	key          int
	value        interface{}
	left         *Node
	right        *Node
	color        bool
}

type RBTree struct {
	sync.RWMutex
	count  int
	root   *Node
}

func newNode(key int, value interface{}, color bool) *Node {
	return &Node{
		key: key,
		value: value,
		color: color,
	}
}

func newRBTree() *RBTree {
	return &RBTree{
		count: 0,
		root:nil,
	}
}

func (self *Node)isRed() bool {
	if self == nil {
		return false
	}
	return self.color
}

func (self *Node) put(key int, val interface{}) *Node {
	if self == nil {
		return newNode(key, val, Red)
	}
	if key < self.key {
		self.left = self.left.put(key, val)
	} else if key > self.key {
		self.right = self.right.put(key, val)
	} else {
		self.value = val
	}
	
	if self.right.isRed() && !self.left.isRed() {
		self = self.rotateLeft()
	}
	if self.left.isRed() && self.left.left.isRed() {
		self = self.rotateRight()
	}
	if self.left.isRed() && self.right.isRed() {
		self.flipColors()
	}
	
	return self
}

func (self *Node) get(key int) interface{} {
	for self != nil {
		if key < self.key {
			self = self.left
		} else if key > self.key {
			self = self.right
		} else {
			return self.value
		}
	}
	return self
}

func (self *Node) hasLeftChild() bool {
	return self.left != nil
}

func (self *Node) hasRightChild() bool {
	return self.right != nil
}

func (self *Node) rotateLeft() *Node {
	node := self.right
	self.right = node.left
	node.left = self
	node.color = node.left.color
	node.left.color = Red
	return node
}

func (self *Node) rotateRight() *Node {
	node := self.left
	self.left = node.right
	node.right = self
	node.color = node.right.color
	node.right.color = Red
	return node
}

func (self *Node) flipColors() {
	self.color = !self.color
	self.left.color = !self.left.color
	self.right.color = !self.right.color
}

func (self *Node) deleteMin() *Node {
	if (self.left == nil) {
		return nil
	}
	if (!self.left.isRed() && !self.left.left.isRed()) {
		self = self.moveRedLeft()
	}
	self.left = self.left.deleteMin()
	return self.balance()
}

func (self *Node) deleteMax() *Node {
	if self.left.isRed() {
		self = self.rotateRight()
	}
	if self.right == nil {
		return nil
	}
	if !self.right.isRed() && !self.right.left.isRed() {
		self = self.moveRedRight()
	}
	self.right = self.right.deleteMax()
	return self.balance()
}

func (self *Node) delete(key int) *Node {
	if key < self.key {
		if !self.left.isRed() && !self.left.left.isRed() {
			self = self.moveRedLeft()
		}
		self.left = self.left.delete(key)
	} else {
		if self.left.isRed() {
			self = self.rotateRight()
		}
		if (key == self.key) && (self.right == nil) {
			return nil
		}
		if !self.right.isRed() && self.right.left.isRed() {
			self = self.moveRedRight()
		}
		if key == self.key {
			node := self.right.min()
			self.key = node.key
			self.value = node.value
			self.right = self.right.deleteMin()
		} else {
			self.right = self.right.delete(key)
		}
	}
	return self.balance()
}

func (self *Node) min() *Node {
	if (self.left == nil) {
		return self
	} else {
		return self.left.min()
	}
}

func (self *Node) balance() *Node {
	if self.right.isRed() {
		self = self.rotateLeft()
	}
	if self.left.isRed() && self.left.left.isRed() {
		self = self.rotateRight()
	}
	if self.left.isRed() && self.right.isRed() {
		self.flipColors()
	}
	return self
}

func (self *Node) moveRedLeft() *Node {
	self.flipColors()
	if self.right.left.isRed() {
		self.right = self.right.rotateRight()
		self = self.rotateLeft()
	}
	return self
}

func (self *Node) moveRedRight() *Node {
	self.flipColors()
	if self.left.left.isRed() {
		self = self.rotateRight()
	}
	return self
}

func (self *RBTree) isEmpty() bool {
	return self.root == nil
}

func (self *RBTree) put(key int, val interface{}) {
	self.root = self.root.put(key, val)
	self.root.color = Black
}

func (self *RBTree) get(key int) interface{} {
	return self.root.get(key)
}

func (self *RBTree) deleteMin() {
	if self.isEmpty() {
		fmt.Println("BST underflow")
		return
	}
	if !self.root.left.isRed() && !self.root.right.isRed() {
		self.root.color = Red
	}
	self.root = self.root.deleteMin()
	if !self.isEmpty() {
		self.root.color = Black
	}
}

func (self *RBTree) deleteMax() {
	if self.isEmpty() {
		fmt.Println("BST underflow")
		return
	}
	if !self.root.left.isRed() && !self.root.right.isRed() {
		self.root.color = Red
	}
	self.root = self.root.deleteMax()
	if !self.isEmpty() {
		self.root.color = Black
	}
}

func (self *RBTree) hasKey(key int) bool {
	return self.get(key) != nil
}

func (self *RBTree) delete(key int) {
	if !self.hasKey(key) {
		fmt.Println("RBTree doesn't has key: ", key)
		return
	}
	if !self.root.left.isRed() && !self.root.right.isRed() {
		self.root.color = Red
	}
	self.root = self.root.delete(key)
	if !self.isEmpty() {
		self.root.color = Black
	}
}

func main() {
	rbtree := newRBTree()
	rbtree.put(1, "hello")
	fmt.Println(rbtree.root.color)
	fmt.Println(rbtree.root.key)
	fmt.Println(rbtree.root.value)
	rbtree.put(2, "world")
	fmt.Println(rbtree.root.value)
	rbtree.put(3, "你好")
	fmt.Println(rbtree.root.value)
	rbtree.put(4, "你好")
	fmt.Println(rbtree.root.value)
	rbtree.put(5, "你好")
	fmt.Println(rbtree.root.value)
	rbtree.put(6, "世界")
	fmt.Println(rbtree.root.value)
	fmt.Println("min", rbtree.root.min())
	hello := rbtree.get(1)
	fmt.Println("get 6 ", hello)
	rbtree.deleteMin()
	fmt.Println(rbtree.root.value)
	rbtree.deleteMax()
	rbtree.deleteMin()
	fmt.Println(rbtree.root.value)
	rbtree.deleteMin()
	fmt.Println(rbtree.root.value)
	rbtree.deleteMin()
	fmt.Println(rbtree.root.value)
	rbtree.deleteMin()
	fmt.Println(rbtree.root)
	rbtree.deleteMin()
	fmt.Println(rbtree.root)
	rbtree.deleteMin()
	fmt.Println(rbtree.root)
}
