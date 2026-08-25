func leastInterval(tasks []byte, n int) int {
	// build freq map
	freqMap := map[byte]int{}
	for _, task := range tasks {
		freq := freqMap[task]
		if freq == 0 {
			freqMap[task] = 1
			continue
		}
		freqMap[task]++		
	}

	// init heap
	heap := NewMaxHeap()
	for name, freq := range freqMap {
		heap.Add(Task{
			name: name,
			freq: freq,
		})
	}

	//do cycle
	cycle := 0
	cdQueue := NewQueueCD(n)
	for heap.Len() > 0 || cdQueue.Len() > 0 {
		cycle++
		hasCD, cd := cdQueue.Peek()
		if hasCD && cd.avail == cycle {
			heap.Add(cd.task)
			cdQueue.Pop()
		}

		hasTask, task := heap.Pop()
		if hasTask {
			task.freq = task.freq-1
			if task.freq > 0 {
				cdQueue.Push(task,cycle)
			}
		}
	}

	return cycle
}

type Cooldown struct {
	task Task
	avail int
}

type QueueCD struct {
	queue []Cooldown
	cooldown int
}

func NewQueueCD(cooldown int) *QueueCD {
	return &QueueCD{
		queue: make([]Cooldown,0),
		cooldown: cooldown,
	}
}

func (q *QueueCD) Len() int {
	return len(q.queue)
}

func (q *QueueCD) Peek() (bool, Cooldown) {
	if q.Len() == 0 {
		return false, Cooldown{}
	}
	return true, q.queue[0]
}

func (q *QueueCD) Pop() Cooldown {
	ret := q.queue[0]
	q.queue = q.queue[1:]
	return ret
}

func (q *QueueCD) Push(t Task, cycle int) {
	q.queue = append(q.queue, Cooldown{
		task: t,
		avail: cycle+q.cooldown+1,
	})
}

type Task struct {
	name byte
	freq int
}

type MaxHeap struct {
	heap []Task
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap: make([]Task,0),
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

func (h *MaxHeap) Add(t Task) {
	h.heap = append(h.heap, t)
	pos := h.Len()-1

	for pos > 0 {
		parent := (pos-1)/2
		if h.heap[parent].freq > h.heap[pos].freq {
			break
		}
		h.Swap(parent,pos)
		pos = parent
	}
}

func (h *MaxHeap) Pop() (bool, Task) {
	if h.Len() == 0 {
		return false, Task{}
	}
	ret := h.heap[0]
	h.Swap(0,h.Len()-1)
	h.heap = h.heap[:h.Len()-1]
	pos := 0
	
	for pos < h.Len() {
		left := 2*pos+1
		right := 2*pos+2
		swap := pos

		if left < h.Len() && h.heap[left].freq > h.heap[swap].freq {
			swap = left
		}
		if right < h.Len() && h.heap[right].freq > h.heap[swap].freq {
			swap = right
		}
		if swap == pos {
			break
		}
		h.Swap(pos,swap)
		pos = swap
	}
	return true, ret
}