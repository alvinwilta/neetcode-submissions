type KthLargest struct {
    k int
	minHeap []int
}


func Constructor(k int, nums []int) KthLargest {
    obj := KthLargest{
		k: k,
		minHeap: make([]int,0),
	}

	for _, num := range nums {
		obj.Add(num)
	}

	return obj
}

func (this *KthLargest) Swap(idxA, idxB int) {
	tmp := this.minHeap[idxA]
	this.minHeap[idxA] = this.minHeap[idxB]
	this.minHeap[idxB] = tmp
}

func (this *KthLargest) bubbleUp(val int) {
	this.minHeap = append(this.minHeap, val)
	// bubble up
	valIdx := len(this.minHeap)-1

	// stop condition: valIdx is the minimal
	for valIdx != 0 {
		// check with parent
		parentIdx := (valIdx-1)/2
		if this.minHeap[valIdx] >= this.minHeap[parentIdx] {
			break
		}

		this.Swap(valIdx, parentIdx)
		valIdx = parentIdx
	}
}

func (this *KthLargest) bubbleDown() {
	// move rightmost to top to shuffle the order
	this.minHeap[0] = this.minHeap[len(this.minHeap)-1]
	this.minHeap = this.minHeap[:len(this.minHeap)-1]
	idx := 0
	length := len(this.minHeap)
	for {
		minimal := idx
		left, right := (2*idx)+1, (2*idx)+2

		if left < length && this.minHeap[left] < this.minHeap[minimal] {
			minimal = left
		}

		if right < length && this.minHeap[right] < this.minHeap[minimal] {
			minimal = right
		}

		if minimal == idx {
			break
		}

		this.Swap(idx, minimal)
		idx = minimal
	}
}


func (this *KthLargest) Add(val int) int {
	this.bubbleUp(val)

	if len(this.minHeap) > this.k {
		this.bubbleDown()
	}
	
	return this.minHeap[0]
}
