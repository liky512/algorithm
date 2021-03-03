package _6_linkedlist

import "testing"

func TestLinkedList_InsertToHead(t *testing.T) {
	 l := NewLinkedList()
	 for i := 0; i < 10;i ++ {
	 	l.InsertToHead(i+1)
	 }
	 l.Print()
}

func TestLinkedList_InsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10;i ++ {
		l.InsertToTail(i+1)
	}
	l.Print()
}

func TestLinkedList_FindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(11))
}

func TestLinkedList_DeleteNode(t *testing.T) {
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
func TestLinkedList_InsertBefore(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10;i ++ {
		l.InsertToTail(i+1)
	}
	l.Print()

	l.InsertBefore(l.head.next.next, 99)
	l.Print()
}