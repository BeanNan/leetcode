package queuebystack

import "errors"

type MyQueue struct {
	pushStack *Stack
	popStack  *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		pushStack: &Stack{},
		popStack:  &Stack{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.pushStack.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.migrate()
	if this.popStack.Empty() {
		return -1
	}
	popValue, _ := this.popStack.Pop()

	return popValue
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.migrate()
	if this.popStack.Empty() {
		return -1
	}
	topValue, _ := this.popStack.Peek()
	return topValue

}

func (this *MyQueue) migrate() bool {
	if !this.popStack.Empty() {
		return false
	}

	if this.pushStack.Empty() {
		return false
	}

	// popStack empty pushStack not empty

	for !this.pushStack.Empty() {
		popValue, _ := this.pushStack.Pop()
		this.popStack.Push(popValue)
	}

	return true

}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.pushStack.Empty() && this.popStack.Empty()
}

type Stack struct {
	arr []int
}

func (s *Stack) Peek() (int, error) {
	if s.Empty() {
		return 0, errors.New("stack empty")
	}

	return s.arr[len(s.arr)-1], nil
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, errors.New("stack empty")
	}

	value := s.arr[len(s.arr)-1]
	s.arr = s.arr[0 : len(s.arr)-1]
	return value, nil

}

func (s *Stack) Push(value int) {
	s.arr = append(s.arr, value)
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Empty() bool {
	return len(s.arr) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
