package chapter6

import "fmt"

type heap struct {
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

func (h *heap) maxHeapify(i int) {
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

func (h *heap) minHeapify(i int) {
	l, r, largest := left(i), right(i), 0
	if l < h.heapSize && h.A[l] < h.A[i] {
		largest = l
	} else {
		largest = i
	}
	if r < h.heapSize && h.A[r] < h.A[largest] {
		largest = r
	}
	if largest != i {
		h.A[i], h.A[largest] = h.A[largest], h.A[i]
		h.minHeapify(largest)
	}
}

func BuildMaxHeap(A []int) heap {
	h := heap{}
	h.A, h.heapSize = A, len(A)
	for i := h.heapSize / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
	return h
}

func BuildMinHeap(A []int) heap {
	h := heap{}
	h.A, h.heapSize = A, len(A)
	for i := h.heapSize / 2; i >= 0; i-- {
		h.minHeapify(i)
	}
	return h
}

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
func (h *heap) InsertMax(x int) {
	h.heapSize++
	h.A[h.heapSize-1] = x - 1
	h.IncreaseKey(h.heapSize-1, x)
}

func (h *heap) Maximum() int {
	return h.A[0]
}

func (h *heap) ExtractMax() (int, error) {
	if h.heapSize < 1 {
		return 0, fmt.Errorf("heap overflow")
	}
	max := h.A[0]
	h.A[0] = h.A[h.heapSize-1]
	h.heapSize--
	h.maxHeapify(0)
	return max, nil
}

func (h *heap) IncreaseKey(i int, key int) error {
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
func (h *heap) Minimum() int {
	return h.A[0]
}

func (h *heap) ExtractMin() (int, error) {
	if h.heapSize < 1 {
		return 0, fmt.Errorf("heap overflow")
	}
	min := h.A[0]
	h.A[0] = h.A[h.heapSize-1]
	h.heapSize--
	h.minHeapify(0)
	return min, nil
}

func (h *heap) DecreaseKey(i int, key int) error {
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

func (h *heap) InsertMin(x int) {
	h.heapSize++
	h.A[h.heapSize-1] = x + 1
	h.DecreaseKey(h.heapSize-1, x)
}
