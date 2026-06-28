type MinStack struct {
	stack []int
	minValMap map[int]int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int,0),
		minValMap: make(map[int]int),
	}
}

func (this *MinStack) Push(val int) {
	length := len(this.stack)
	if length == 0 {
		this.minValMap[length] = val
	} else {
		if val < this.minValMap[length-1] {
			this.minValMap[length] = val
		} else {
			this.minValMap[length] = this.minValMap[length-1]
		}
	}

	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minValMap[len(this.stack)-1]
}
