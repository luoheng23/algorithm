package chapter16

// Node for code tree
type Node struct {
	value       rune
	freq        int
	left, right *Node
}

// CodeTree is tree for code
type CodeTree struct {
	root *Node
}

// Heap represents heap data strcuture
type Heap struct {
	A        []*Node
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

func (h *Heap) minHeapify(i int) {
	l, r, smallest := left(i), right(i), 0
	if l < h.heapSize && h.A[l].freq < h.A[i].freq {
		smallest = l
	} else {
		smallest = i
	}
	if r < h.heapSize && h.A[r].freq < h.A[smallest].freq {
		smallest = r
	}
	if smallest != i {
		h.A[i], h.A[smallest] = h.A[smallest], h.A[i]
		h.minHeapify(smallest)
	}
}

// BuildMinHeap O(n)
func BuildMinHeap(A []*Node) Heap {
	h := Heap{}
	h.A, h.heapSize = A, len(A)
	for i := h.heapSize / 2; i >= 0; i-- {
		h.minHeapify(i)
	}
	return h
}

// minimum priority queue
// method:
// Minimum
// ExtractMin
// DecreaseKey
// InsertMin

// Minimum O(1)
func (h *Heap) Minimum() *Node {
	return h.A[0]
}

// ExtractMin O(lgn)
func (h *Heap) ExtractMin() *Node {
	if h.heapSize < 1 {
		return nil
	}
	min := h.A[0]
	h.A[0] = h.A[h.heapSize-1]
	h.heapSize--
	h.minHeapify(0)
	return min
}

// DecreaseKey O(lgn)
func (h *Heap) DecreaseKey(i int, freq int) error {
	h.A[i].freq = freq
	for i > 0 && h.A[parent(i)].freq > h.A[i].freq {
		h.A[i], h.A[parent(i)] = h.A[parent(i)], h.A[i]
		i = parent(i)
	}
	return nil
}

// InsertMin O(lgn)
func (h *Heap) InsertMin(x *Node) {
	h.heapSize++
	h.A[h.heapSize-1] = x
	h.A[h.heapSize-1].freq--
	h.DecreaseKey(h.heapSize-1, x.freq+1)
}

// Huffman produce a code Tree
func Huffman(c []*Node) CodeTree{
	n := len(c)
	q := BuildMinHeap(c)
	for i:= 0; i < n - 1; i++ {
		z := Node{}
		z.left = q.ExtractMin()
		z.right = q.ExtractMin()
		z.freq = z.left.freq + z.right.freq
		q.InsertMin(&z)
	}
	codeTree := CodeTree{root: (q.ExtractMin())}
	return codeTree
}