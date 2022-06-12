// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)

	// heapify up from last index of array
	h.maxHeapifyUp(len(h.array) - 1)
}

// maxHeapifyUp - takes an integer indicating where the heapify is going to start
func (h *MaxHeap) maxHeapifyUp(index int) {
	// swap when the current index is larger than the parent

	// when the parent is smaller than the current index
	for h.array[parent(index)] < h.array[index] {
		// then swap
		h.swap(parent(index), index) // swaps until it finds its place

		// after swapping perform the same on the parent ( which is the swapped key )
		index = parent(index) // updates its current index to its parent index so that we can loop again
	}
}

// maxHeapifyDown will heapify top to bottom
// 1. get the larger child
// 2. compare to the current index
// 3. swap it
// 4. repeat loop until index has no children
func (h *MaxHeap) maxHeapifyDown(index int) {
	// if the left child index is the same or smaller than the last of the array - you can say this index has at least 1 child
	lastIndex := len(h.array) - 1
	lft, rgt := left(index), right(index)

	// childToCompare is child to compare with the index and to swap
	childToCompare := 0

	// loop while index has at least one child
	for lft <= lastIndex {
		// if there is only one child
		// or the left side is larger
		// assign left child as child to compare

		if lft == lastIndex { // when left child is the only child
			childToCompare = lft
		} else if h.array[lft] > h.array[rgt] { // when left child is larger
			childToCompare = lft
		} else { // when right child is larger
			childToCompare = rgt
		}

		// compare array value of current index to larger child
		// and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			// update vars in the loop so next round this will
			// work on the swapped index
			index = childToCompare
			lft, rgt = left(index), right(index)
		} else { // if the value of the current index is larger than the largest child
			return
		}
	}
}

func parent(i int) int {
	// the parent index can be expressed as index -1 divided by 2
	// left child is always an odd number | right child always an even #
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the right child index
func right(i int) int {
	return 2*1 + 2
}

// swap keys in array (will mutate array)
func (h *MaxHeap) swap(i1, i2 int) {
	// swapping putting i2 into i1 && i1 into i2
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// Extract returs the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	lastIndex := len(h.array) - 1

	// don't extract from empty array
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	// get the last index and put in the root (index[0])
	h.array[0] = h.array[lastIndex]
	// shorten the array
	h.array = h.array[:lastIndex]

	h.maxHeapifyDown(0)
	return extracted
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 14, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	// extract 5x's
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
