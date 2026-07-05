package medium

type MinStack struct {
	stack       []int
	stackOfMins []int
}

func ConstructorMS() MinStack {
	return MinStack{
		stack:       []int{},
		stackOfMins: []int{},
	}
}

func (this *MinStack) Push(value int) {
	if len(this.stack) == 0 || value <= this.stackOfMins[len(this.stackOfMins)-1] {
		this.stackOfMins = append(this.stackOfMins, value)
	}
	this.stack = append(this.stack, value)
}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack)-1] == this.stackOfMins[len(this.stackOfMins)-1] {
		this.stackOfMins = this.stackOfMins[:len(this.stackOfMins)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.stackOfMins[len(this.stackOfMins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
