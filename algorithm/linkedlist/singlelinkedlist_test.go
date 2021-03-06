package linkedlist

import "testing"

var l *LinkedList

func init() {
	n5 := &ListNode{value: 5}
	n4 := &ListNode{value: 4, next: n5}
	n3 := &ListNode{value: 3, next: n4}
	n2 := &ListNode{value: 2, next: n3}
	n1 := &ListNode{value: 1, next: n2}
	l = &LinkedList{head: &ListNode{next: n1}}
}

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i<10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i<10; i++ {
		l.InsertToTail(i + 1)
	}

	l.Print()
}

func testFindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i:=0; i<10; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.head.next))
	l.Print()

	t.Log(l.DeleteNode(l.head.next.next))
	l.Print()
}

func TestLinkedList_Reverse(t *testing.T) {
	l.Print()
	l.Reverse()
	l.Print()
}

func TestLinkedList_HasCycle(t *testing.T) {
	t.Log(l.HasCycle())
	l.head.next.next.next.next.next.next = l.head.next.next.next
	t.Log(l.HasCycle())
}

func TestMergeSortedList(t *testing.T) {
	n5 := &ListNode{value: 9}
	n4 := &ListNode{n5, 7}
	n3 := &ListNode{n4, 5}
	n2 := &ListNode{n3, 3}
	n1 := &ListNode{n2, 1}
	l1 := &LinkedList{head: &ListNode{next: n1}}

	n10 := &ListNode{value: 10}
	n9 := &ListNode{value: 8, next: n10}
	n8 := &ListNode{value: 6, next: n9}
	n7 := &ListNode{value: 4, next: n8}
	n6 := &ListNode{value: 2, next: n7}
	l2 := &LinkedList{head: &ListNode{next: n6}}

	MergeSortedList(l1, l2).Print()

}

//func TestDeleteBottomN(t *testing.T) {
//	l.Print()
//	//l.DeleteBottomN(3)
//	l.Print()
//}
//
//func TestFindMiddleNode(t *testing.T) {
//	l.DeleteBottomN(1)
//	l.DeleteBottomN(1)
//	l.DeleteBottomN(1)
//	l.DeleteBottomN(1)
//	l.Print()
//	t.Log(l.FindMiddleNode())
//}




