/*

1. It keeps the items in a min-heap data structure.
2. It swaps the root with the right most child node, removes it from the heap, and insert it in the final array.
3. Then it moves the smallest remaining item to the root.
4. It stops when the root is empty.

*/

package main

type MaxHeap struct {
    slice    []int
    heapSize int
}

func (h *MaxHeap) heapify(length, i int) {
    largest := i
    left := 2*i + 1
    right := 2*i + 2

    if left < length && h.slice[left] > h.slice[largest] {
        largest = left
    }

    if right < length && h.slice[right] > h.slice[largest] {
        largest = right
    }

    if largest != i {
        h.slice[i], h.slice[largest] = h.slice[largest], h.slice[i]
        h.heapify(length, largest)
    }
}

func (h *MaxHeap) sort() {
    length := len(h.slice)

    for i := length/2 - 1; i >= 0; i-- {
        h.heapify(length, i)
    }

    for i := length - 1; i >= 0; i-- {
        h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
        h.heapify(i, 0)
    }
}
