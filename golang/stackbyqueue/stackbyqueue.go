package stackbyqueue

type MyStack struct {
	pushQueue *Queue
	popQueue  *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		pushQueue: &Queue{},
		popQueue:  &Queue{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.pushQueue.Push(x)

	for !this.popQueue.Empty() {
		this.pushQueue.Push(this.popQueue.Pop())
	}
	this.pushQueue, this.popQueue = this.popQueue, this.pushQueue

}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.popQueue.Empty() {
		return -1
	}
	popValue := this.popQueue.Pop()

	return popValue
}

/** Get the top element. */
func (this *MyStack) Top() int {

	if this.popQueue.Empty() {
		return -1
	}

	return this.popQueue.Peek()

}



/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.pushQueue.Empty() && this.popQueue.Empty()
}

type Queue struct {
	arr []int
}

func (q *Queue) Push(x int) {
	q.arr = append(q.arr, x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *Queue) Pop() int {
	if q.Empty() {
		return -1
	}

	popValue := q.arr[0]
	q.arr = q.arr[1:]
	return popValue

}

/** Get the front element. */
func (q *Queue) Peek() int {
	if q.Empty() {
		return -1
	}

	return q.arr[0]
}

/** Returns whether the queue is empty. */
func (q *Queue) Empty() bool {
	return len(q.arr) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
