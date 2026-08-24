func kClosest(points [][]int, k int) [][]int {
    maxHeap := NewMaxHeap(k)
    for i, point := range points {
        dist := CalcDist(point)
        maxHeap.Add(CalcRes{
            dist: dist,
            idx: i,
        })

		if maxHeap.Len() > k {
			maxHeap.Pop()
		}
    }

    // get the solution
    res := [][]int{}
    for range k {
        res = append(res, points[maxHeap.Pop().idx])
    }

    return res
}

func CalcDist(points []int) float64 {
    return math.Sqrt(math.Pow(float64(points[0]),2)+math.Pow(float64(points[1]),2))
}

type MaxHeap struct {
    k int
    heap []CalcRes 
}

type CalcRes struct {
    dist float64 // the distance
    idx int // the index of points 
}

func NewMaxHeap(k int) *MaxHeap {
    return &MaxHeap{
        k: k,
        heap: make([]CalcRes,0),
    }
}

func (h *MaxHeap) Len() int {
    return len(h.heap)
}

func (h *MaxHeap) GetMaxDist() float64 {
    return h.heap[0].dist
}

func (h *MaxHeap) Swap(a,b int) {
    tmp := h.heap[a]
    h.heap[a] = h.heap[b]
    h.heap[b] = tmp
}


func (h *MaxHeap) Add(point CalcRes) {
    h.heap = append(h.heap, point)
    pos := h.Len()-1

    for pos != 0 {
        parent := (pos-1)/2
        if h.heap[parent].dist < h.heap[pos].dist {
            h.Swap(parent, pos)
            pos = parent
            continue
        }
        break
    }
}

func (h *MaxHeap) Pop() CalcRes {
    if h.Len() == 0 {
        return CalcRes{}
    }
    // swap and remove
    res := h.heap[0]
    h.Swap(0,h.Len()-1)
    h.heap = h.heap[:h.Len()-1]
    pos := 0

    // bubble down
    for pos != h.Len()-1 {
        left := (2*pos)+1
        right := (2*pos)+2
        swap := pos
        if left < h.Len() && h.heap[left].dist > h.heap[swap].dist {
            swap = left
        }
        if right < h.Len() && h.heap[right].dist > h.heap[swap].dist {
            swap = right
        }

        if swap == pos {
            break
        }
        h.Swap(pos,swap)
        pos = swap
    }

    return res
}