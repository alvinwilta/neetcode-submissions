func findKthLargest(nums []int, k int) int {
	heap := NewMinHeap(k)
	for _, num := range nums {
		if heap.Len() == k && heap.GetMin() > num {
			continue
		}
		heap.Add(num)

		if heap.Len() > k {
			heap.Pop()
		}
	}

	return heap.GetMin()
}

type MinHeap struct {
	heap []int
	k int
}

func NewMinHeap(k int) *MinHeap {
	return &MinHeap{
		heap: make([]int,0),
		k: k,
	}
}

func (h *MinHeap) Len() int {
	return len(h.heap)
}

func (h *MinHeap) GetMin() int {
	return h.heap[0]
}

func (h *MinHeap) Swap(a,b int) {
	tmp := h.heap[a]
	h.heap[a] = h.heap[b]
	h.heap[b] = tmp
}

func (h *MinHeap) Add(num int) {
	h.heap = append(h.heap,num)
	pos := h.Len()-1

	for pos > 0 {
		parent := (pos-1)/2
		if h.heap[parent] < h.heap[pos] {
			break
		}

		h.Swap(parent,pos)
		pos = parent
	}
}

func (h *MinHeap) Pop() int {
	if h.Len() == 0 {
		return 0
	}
	res := h.heap[0] // store root as result
	h.Swap(0,h.Len()-1) // swap root
	h.heap = h.heap[:h.Len()-1] // remove root
	root := 0

	for root < h.Len() {
		left := 2*root+1
		right := 2*root+2
		swap := root

		if left < h.Len() && h.heap[left] < h.heap[swap] {
			swap = left
		}

		if right < h.Len() && h.heap[right] < h.heap[swap] {
			swap = right
		}

		if swap == root {
			break
		}
		h.Swap(root,swap)
		root = swap
	}

	return res
}