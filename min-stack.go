type MinStack struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minStack := new(MinStack)
	return *minStack
}

func (this *MinStack) Push(val int) {
	if len(this.stack) < 1 {
		this.min = val
	} else {
		if val < this.min {
			this.min = val
		}
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) < 1 {
		return
	}
	popped := this.Top()
	this.stack = this.stack[:len(this.stack)-1]
	if popped == this.min && len(this.stack) > 0 {
		this.min = this.stack[0]

		for _, n := range this.stack {
			if n < this.min {
				this.min = n
			}
		}
	}

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// TO1

type MinStack struct {
	Stack    []int
	MinStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(value int) {
	if len(this.MinStack) < 1 || value < this.MinStack[len(this.MinStack)-1] {
		this.MinStack = append(this.MinStack, value)
	} else {
		this.MinStack = append(this.MinStack, this.GetMin())
	}

	this.Stack = append(this.Stack, value)
}

func (this *MinStack) Pop() {
	this.MinStack = this.MinStack[:len(this.MinStack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinStack[len(this.MinStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */