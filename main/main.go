package main

import (
	"fmt"
	"main/datastructure/hashtable"
	"main/datastructure/linkedlist"
)

func linkedListTest() {
	l := linkedlist.InitList()
	l.PushFront("1")
	l.PushFront("2")
	l.PushFront("3")
	l.PushBack("10")
	l.Print()
	l.RemoveFront()
	l.Print()
	l.RemoveBack()
	l.Print()
}

func main() {
	ht := hashtable.InitHashTable()
	ht.Set("1", "ana")
	ht.Set("2", "popescu")
	ht.Set("3", "georgescu")

	fmt.Println(ht.Get("3"))
}
