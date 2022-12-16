package main

func main() {
	head := []int{1, 2, 3, 4, 5}

	rotateRight(head, 2)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	iter := head
	//这个是链表不是数组，结构体中并没有定义链表的长度所以应该用循环来计算出链表的长度
	//curr为指针，指针不断向后移动，最后指针在链表尾部停下俩，让指针的next指向head就可以实现链表的循环
	//1->2->3->4->1->2->3->4
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	//注意此时的指针还在head的尾部
	add := n - k%n
	//计算出链表的偏移量
	if add == n {
		return head
	}
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}
