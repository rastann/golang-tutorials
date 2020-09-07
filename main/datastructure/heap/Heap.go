package heap

type Heap struct {
	Items []int
}

func (h *Heap) GetLeftIndex(parentIndex int) int {
	return 2*parentIndex+1
}

func (h *Heap) GetRightIndex(parentIndex int) int {
	return 2*parentIndex+2
}

func (h *Heap) HasLeft(index int) bool {
	return h.GetLeftIndex(index) < len(h.Items)
}

func (h *Heap) HasRight(index int) bool {
	return h.GetRightIndex(index) < len(h.Items)
}

func (h *Heap) GetLeftChild(parentIndex int) int {
	return h.Items[h.GetLeftIndex(parentIndex)]
}

func (h *Heap) GetRightChild(parentIndex int) int {
	return h.Items[h.GetRightIndex(parentIndex)]
}

func (h *Heap) Swap(indexOne, indexTwo int) {
	h.Items[indexOne], h.Items[indexTwo] = h.Items[indexTwo], h.Items[indexOne]
}

func (h *Heap) GetParent(index int) int {
	return h.Items[h.GetParentIndex(index)]
}

func (h *Heap) GetParentIndex(childIndex int) int {
return (childIndex - 1) / 2
}

func (h *Heap) HasParent(index int) bool {
	return h.GetParentIndex(index) >= 0
}