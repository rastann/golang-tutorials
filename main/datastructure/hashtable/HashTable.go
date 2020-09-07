package hashtable

import (
	"container/list"
	"math"
)

const arrayLength = 100

type HashTable struct {
	data[arrayLength] *list.List
}

type listData struct {
	key string
	value interface{}
}

func InitHashTable() *HashTable {
	return &HashTable{[arrayLength]*list.List{} }
}

func hash(s string) int {
	h := 0
	for pos, char := range s {
		h += int(char) * int(math.Pow(31, float64(len(s)-pos+1)))
	}
	return h
}

func index(hash int) int {
	return hash % arrayLength
}

func (h *HashTable) Set(k string, v interface{}) *HashTable {
	index := index(hash(k))

	if h.data[index] == nil {
		h.data[index] = list.New()
		h.data[index].PushBack(listData{k,v})
	} else {
		node := h.data[index].Front()
		for {
			if node != nil {
				d := node.Value.(listData)
				if d.key == k {
					d.value = v
					break
				}
			} else {
				h.data[index].PushBack(listData{k,v})
				break
			}
			node = node.Next()
		}
	}
	return h
}

func (h *HashTable) Get(k string) (result interface{}, ok bool) {
	index := index(hash(k))
	linkedList := h.data[index]

	if linkedList == nil {
		return "", false
	}
	node := linkedList.Front()
	for {
		if node != nil {
			d := node.Value.(listData)
			if d.key == k {
				return d.value, true
			}
		} else {
			return "", false
		}
		node = node.Next()
	}
}
