package linkedlist

import (
	"fmt"
)

type node struct {
	val  string
	next *node
}

type List struct {
	len int
	head *node
}

func InitList() *List {
	return &List{}
}

func (l *List) Print() error {
	if l.head == nil {
		return fmt.Errorf("Print: List is empty")
	}
	c := l.head
	for c != nil {
		fmt.Print(c.val, " ")
		c = c.next
	}
	fmt.Println()
	return nil
}

func (l *List) PushFront(v string) {
	n := &node{val: v}
	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.len++
}

func (l *List) PushBack(v string) {
	n := &node{val: v}
	if l.head == nil {
		l.head = n
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}
	l.len++
}

func (l *List) RemoveFront() error {
	if l.head == nil {
		return fmt.Errorf("RemoveFront: List is empty")
	}
	l.head = l.head.next
	l.len--
	return nil
}

func (l *List) RemoveBack() error {
	if l.head == nil {
		return fmt.Errorf("RemoveBack: List is empty")
	}
	var prev *node
	current := l.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		l.head = nil
	}
	l.len--
	return nil
}

func (l *List) Size() int {
	return l.len
}

func (l *List) Pop() (string, error) {
	if l.head == nil {
		return "", fmt.Errorf("Pop: List is empty")
	}
	return l.head.val, nil
}