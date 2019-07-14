package chapter6

import "fmt"

// Heap represents heap data strcuture
type Heap struct {
	A        []int
	heapSize int
}

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func (h *Heap) maxHeapify(i int) {
	l, r, largest := left(i), right(i), 0
	if l < h.heapSize && h.A[l] > h.A[i] {
		largest = l
	} else {
		largest = i
	}
	if r < h.heapSize && h.A[r] > h.A[largest] {
		largest = r
	}
	if largest != i {
		h.A[i], h.A[largest] = h.A[largest], h.A[i]
		h.maxHeapify(largest)
	}
}

func (h *Heap) minHeapify(i int) {
	l, r, smallest := left(i), right(i), 0
	if l < h.heapSize && h.A[l] < h.A[i] {
		smallest = l
	} else {
		smallest = i
	}
	if r < h.heapSize && h.A[r] < h.A[smallest] {
		smallest = r
	}
	if smallest != i {
		h.A[i], h.A[smallest] = h.A[smallest], h.A[i]
		h.minHeapify(smallest)
	}
}

// BuildMaxHeap O(n)
func BuildMaxHeap(A []int) Heap {
	h := Heap{}
	h.A, h.heapSize = A, len(A)
	for i := h.heapSize / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
	return h
}

// BuildMinHeap O(n)
func BuildMinHeap(A []int) Heap {
	h := Heap{}
	h.A, h.heapSize = A, len(A)
	for i := h.heapSize / 2; i >= 0; i-- {
		h.minHeapify(i)
	}
	return h
}

// HeapSort O(nlgn)
func HeapSort(A []int) {
	h := BuildMaxHeap(A)
	for i := len(A) - 1; i >= 0; i-- {
		h.A[0], h.A[i] = h.A[i], h.A[0]
		h.heapSize--
		h.maxHeapify(0)
	}
}

// maximum priority queue
// method:
// InsertMax
// Maximum
// ExtractMax
// IncreaseKey

// InsertMax O(lgn)
func (h *Heap) InsertMax(x int) {
	h.heapSize++
	h.A[h.heapSize-1] = x - 1
	h.IncreaseKey(h.heapSize-1, x)
}

// Maximum O(1)
func (h *Heap) Maximum() int {
	return h.A[0]
}

// ExtractMax O(lgn)
func (h *Heap) ExtractMax() (int, error) {
	if h.heapSize < 1 {
		return 0, fmt.Errorf("Heap overflow")
	}
	max := h.A[0]
	h.A[0] = h.A[h.heapSize-1]
	h.heapSize--
	h.maxHeapify(0)
	return max, nil
}

// IncreaseKey O(lgn)
func (h *Heap) IncreaseKey(i int, key int) error {
	if key < h.A[i] {
		return fmt.Errorf("new key is smaller than current key")
	}
	h.A[i] = key
	for i > 0 && h.A[parent(i)] < h.A[i] {
		h.A[i], h.A[parent(i)] = h.A[parent(i)], h.A[i]
		i = parent(i)
	}
	return nil
}

// minimum priority queue
// method:
// Minimum
// ExtractMin
// DecreaseKey
// InsertMin

// Minimum O(1)
func (h *Heap) Minimum() int {
	return h.A[0]
}

// ExtractMin O(lgn)
func (h *Heap) ExtractMin() (int, error) {
	if h.heapSize < 1 {
		return 0, fmt.Errorf("Heap overflow")
	}
	min := h.A[0]
	h.A[0] = h.A[h.heapSize-1]
	h.heapSize--
	h.minHeapify(0)
	return min, nil
}

// DecreaseKey O(lgn)
func (h *Heap) DecreaseKey(i int, key int) error {
	if key > h.A[i] {
		return fmt.Errorf("new key is bigger than current key")
	}
	h.A[i] = key
	for i > 0 && h.A[parent(i)] > h.A[i] {
		h.A[i], h.A[parent(i)] = h.A[parent(i)], h.A[i]
		i = parent(i)
	}
	return nil
}

// InsertMin O(lgn)
func (h *Heap) InsertMin(x int) {
	h.heapSize++
	h.A[h.heapSize-1] = x + 1
	h.DecreaseKey(h.heapSize-1, x)
}
