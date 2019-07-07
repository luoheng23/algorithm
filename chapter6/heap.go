package chapter6

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
