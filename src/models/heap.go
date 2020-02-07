package models

import (
	"../utils"
)

// Heap is the structure of the heap data structure
type Heap struct {
	array           []int
	count, capacity int
	cmpFunc         func(int, int) bool
}

// IsEmpty return true if the heap is empty
func (h *Heap) IsEmpty() bool { return h.count == 0 }

// GetCapacity return the capacity value of the heap
func (h *Heap) GetCapacity() int { return h.capacity }

type heapType int

const (
	// HeapMin is used to set the heap min type
	HeapMin heapType = 0
	// HeapMax is used to set the heap max type
	HeapMax heapType = 1
)

// New creates a new heap structure
func (h *Heap) New(capacity int, kind heapType) *Heap {
	h.array = make([]int, capacity)
	h.capacity = capacity
	h.count = 0
	if kind == HeapMin {
		h.cmpFunc = minOrder
	} else if kind == HeapMax {
		h.cmpFunc = maxOrder
	}
	return h
}

// Insert add a new value in the heap
func (h *Heap) Insert(value int) {
	if h.count >= h.capacity {
		return
	}
	h.array[h.count] = value
	h.heapify(h.count, h.cmpFunc)
	h.count++
}

func minOrder(a, b int) bool {
	return a < b
}

func maxOrder(a, b int) bool {
	return a > b
}

func (h *Heap) heapify(nodeIndex int, cmp func(int, int) bool) {
	parentNodeIndex := (nodeIndex - 1) / 2
	if cmp(h.array[nodeIndex], h.array[parentNodeIndex]) {
		utils.SwapInt(&h.array[nodeIndex], &h.array[parentNodeIndex])
		h.heapify(parentNodeIndex, cmp)
	}
}

func (h *Heap) pop() {
	h.capacity--
	h.array = h.array[:h.capacity]
	h.count--
}

// DeleteRoot remove the root by swaping the first and last item
func (h *Heap) DeleteRoot() int {
	root := h.array[0]
	lastNodeIndex := h.count - 1
	utils.SwapInt(&h.array[lastNodeIndex], &h.array[0])
	h.pop()
	h.down(0)
	return root
}

func (h *Heap) down(nodeIndex int) {
	leftChild := nodeIndex*2 + 1
	rightChild := nodeIndex*2 + 2
	extremity := nodeIndex
	// Stop if no childs (if no left child then no right child)
	if leftChild > h.count-1 || leftChild < 0 {
		return
	}
	if h.cmpFunc(h.array[leftChild], h.array[nodeIndex]) {
		extremity = leftChild
	}
	if rightChild <= h.count-1 && rightChild >= 0 &&
		h.cmpFunc(h.array[rightChild], h.array[extremity]) {
		extremity = rightChild
	}
	if extremity != nodeIndex {
		utils.SwapInt(&h.array[extremity], &h.array[nodeIndex])
		h.down(extremity)
	}
	// fmt.Printf("Parent: %d, Left: %d, Right: %d\n", h.array[nodeIndex], h.array[leftChild], h.array[rightChild])
}
