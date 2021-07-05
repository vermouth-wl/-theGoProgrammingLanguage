package main

import "fmt"

// 链表

// Node 使用Struct定义单链表
type Node struct {
	data int
	next *Node
}

// ShowNode 遍历节点
func ShowNode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		// 移动指针
		p = p.next
	}
}

func main() {
	var head = new(Node)
	head.data = 1
	var node1 = new(Node)
	node1.data = 2

	head.next = node1
	var node2 = new(Node)
	node2.data = 3

	node1.next = node2
	var node3 = new(Node)
	node3.data = 4

	node2.next = node3
	ShowNode(head)
}
