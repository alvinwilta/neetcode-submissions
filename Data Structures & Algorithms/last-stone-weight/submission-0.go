func lastStoneWeight(stones []int) int {
	heap := NewMaxHeap()

	for _, stone := range stones {
		heap.Push(stone)
	}

	for heap.Len() > 1 {
		stone1 := heap.Pop()
		stone2 := heap.Pop()
		resStone := max(stone1,stone2) - min(stone1,stone2)
		heap.Push(resStone)
	}

	return heap.Pop()
}

type MaxHeap struct {
	heap []int
}

func NewMaxHeap() MaxHeap {
	return MaxHeap{
		heap: make([]int,0),
	}
}

func (h *MaxHeap) Len() int {
	return len(h.heap)
}

func (h *MaxHeap) Swap(a,b int) {
	tmp := h.heap[a]
	h.heap[a] = h.heap[b]
	h.heap[b] = tmp
}

func (h *MaxHeap) Push(num int) {
	h.heap = append(h.heap, num)
	idx := len(h.heap)-1

	// bubble up
	for idx > 0 {
		parent := (idx-1)/2
		if h.heap[parent] >= h.heap[idx] {
			break
		} 
		
		h.Swap(parent,idx)
		idx = parent
	}
}

func (h *MaxHeap) Pop() int {
	if len(h.heap) == 0 {
		return 0
	}

	ret := h.heap[0]
	h.Swap(0,h.Len()-1)
	h.heap = h.heap[:h.Len()-1]
	idx := 0
	length := h.Len()

	for {
		maxIdx := idx
		left := 2*idx+1
		right := 2*idx+2

		if left < length && h.heap[left] > h.heap[maxIdx] {
			maxIdx = left
		}

		if right < length && h.heap[right] > h.heap[maxIdx] {
			maxIdx = right
		}

		if maxIdx == idx {
			break
		}
		h.Swap(maxIdx,idx)
		idx = maxIdx
	}

	return ret
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}