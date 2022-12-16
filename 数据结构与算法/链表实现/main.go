package main

import "fmt"

type Node struct {
	data interface{}
	pre  *Node
	next *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}


func main() {
	dl := initList()
	fmt.Println("从开头添加节点")
	for i := 0; i < 5; i++ {
		node := Node{
			data: i,
		}
		dl.addNodeBeforeHead(&node)
	}
	dl.printList()
}

func (l *List) getSize() int {
	return l.size
}

//获取链表尾部
func (l *List) getHead() *Node {
	return l.head
}

//获取链表尾部
func (l *List) getTail() *Node {
	return l.tail
}

//初始化链表
func initList() *List {
	return &List{
		head: nil,
		tail: nil,
		size: 0,
	}
}

//打印链表
func (l *List) printList() {
	fmt.Println("链表的长度是:", l.size)
	//定义指针指向头部
	pointer := l.head
	//当指针的next为空时即退出循环,也就是pointer指向不是nil时继续循环
	for pointer != nil {
		fmt.Println("当前节点的值为", pointer.data)
		pointer = pointer.next
	}
}

//在头部前添加节点
func (l *List) addNodeBeforeHead(node *Node) {
	if l.getSize() == 0 {
		l.head = node
		l.tail = node
		node.pre = nil
		node.next = nil
	} else {
		node.pre, node.next, l.head, l.head.pre = l.head.pre, l.head, node, node
	}
	l.size += 1
}

//在尾部后添加节点
func (l *List) addNodeAfterTail(node *Node) {
	if l.getSize() == 0 {
		l.head = node
		l.tail = node
		node.pre = nil
		node.next = nil
	} else {
		node.next = nil
		node.pre = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.size += 1
}
