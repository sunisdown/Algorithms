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

func newNode(key int, value interface{}, color bool) *Node {
	return &Node{
		key: key,
		value: value,
		color: color,
	}
}

func isRed(node *Node) bool {
	if node == nil {
		return false
	}
	return node.color
}

func put(node *Node, key int, val interface{}) *Node {
	if node == nil {
		return newNode(key, val, Red)
	}
	if key < node.key {
		node.left = put(node.left, key, val)
	} else if key > node.key {
		node.right = put(node.right, key, val)
	} else {
		node.value = val
	}
	
	if isRed(node.right) && !isRed(node.left) {
		node = node.rotateLeft()
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = node.rotateRight()
	}
	if isRed(node.left) && isRed(node.right) {
		node.flipColors()
	}
	
	return node
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
	if (!isRed(self.left) && !isRed(self.left.left)) {
		self = self.moveRedLeft()
	}
	self.left = self.left.deleteMin()
	return self.balance()
}

func (self *Node) balance() *Node {
	if isRed(self.right) {
		self = self.rotateLeft()
	}
	if isRed(self.left) && isRed(self.left.left) {
		self = self.rotateRight()
	}
	if isRed(self.left) && isRed(self.right) {
		self.flipColors()
	}
	return self
}

func (self *Node) moveRedLeft() *Node {
	self.flipColors()
	if isRed(self.right.left) {
		self.right = self.right.rotateRight()
		self = self.rotateLeft()
	}
	return self
}

func (self *Node) moveRedRight() *Node {
	self.flipColors()
	if isRed(self.left.left) {
		self = self.rotateRight()
	}
	return self
}

type RBTree struct {
	sync.RWMutex
	count  int
	root   *Node
}

func newRBTree() *RBTree {
	return &RBTree{
		count: 0,
		root:nil,
	}
}

func (self *RBTree) isEmpty() bool {
	return self.root == nil
}

func (self *RBTree) put(key int, val interface{}) {
	self.root = put(self.root, key, val)
	self.root.color = Black
}

func (self *RBTree) get(key int) interface{} {
	return self.root.get(key)
}

func (self *RBTree) deleteMin() {
	if self.isEmpty() {
		fmt.Println("BST underflow")
	}
	if isRed(self.root.left) && !isRed(self.root.right) {
		self.root.color = Red
	}
	self.root = self.root.deleteMin()
	if !self.isEmpty() {
		self.root.color = Black
	}
}

func (self *RBTree) deleteMax() {
}

func (self *RBTree) delete(key int) {
}

func (self *RBTree) remove(node *Node) {
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
	hello := rbtree.get(1)
	fmt.Println("get 6 ", hello)
	rbtree.deleteMin()
	fmt.Println(rbtree.root.value)
	rbtree.deleteMin()
	fmt.Println(rbtree.root.value)
	rbtree.deleteMin()
	fmt.Println(rbtree.root.value)
	rbtree.deleteMin()
	fmt.Println(rbtree.root.value)
	rbtree.deleteMin()
	fmt.Println(rbtree.root)
}
