package maxheap

import (
	"main/datastructure/heap"
)

type MaxHeap struct {
	*heap.Heap
}

func New(input []int) *MaxHeap {
	h := &MaxHeap{&heap.Heap{Items: input}}

	if len(h.Items) > 0 {
		h.buildMaxHeap()
	}

	return h
}

func (h *MaxHeap) buildMaxHeap() {
	for i := len(h.Items)/2-1; i>=0; i-- {
		h.maxHeapifyDown(i)
	}
}

func (h *MaxHeap) Insert(item int) *MaxHeap {
	h.Items = append(h.Items, item)
	lastElementIndex := len(h.Items) - 1
	h.maxHeapifyUp(lastElementIndex)

	return h
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.HasParent(index) && h.GetParent(index) < h.Items[index] {
		h.Swap(h.GetParentIndex(index), index)
		index = h.GetParentIndex(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	for (h.HasLeft(index) && h.Items[index] < h.GetLeftChild(index)) ||
		(h.HasRight(index) && h.Items[index] < h.GetRightChild(index)) {
		if(h.HasLeft(index) && h.Items[index] < h.GetLeftChild(index)) &&
			(h.HasRight(index) && h.Items[index] < h.GetRightChild(index)) {
			if h.GetLeftChild(index) > h.GetRightChild(index) {
				h.Swap(index, h.GetLeftIndex(index))
				index = h.GetLeftIndex(index)
			} else {
				h.Swap(index, h.GetRightIndex(index))
				index = h.GetRightIndex(index)
			}
		} else if h.HasLeft(index) && h.Items[index] < h.GetLeftChild(index) {
			h.Swap(index, h.GetLeftIndex(index))
			index = h.GetLeftIndex(index)
		} else {
			h.Swap(index, h.GetRightIndex(index))
			index = h.GetRightIndex(index)
		}
	}
}