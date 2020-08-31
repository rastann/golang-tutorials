package main

import (
	"main/linkedlist"
)

func main() {
	l := linkedlist.InitList()
	l.PushFront("1")
	l.PushFront("2")
	l.PushFront("3")
	l.PushBack("10")
	l.Print()
	l.RemoveFront()
	l.Print()
}
