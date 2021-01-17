package kthlargest

import "errors"

type KthLargest struct {
	data []int
	cap  int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		cap: k,
	}

	for _, num := range nums {
		kth.Push(num)
	}

	return kth
}

func (this *KthLargest) Add(val int) int {
	this.Push(val)

	topValue, _ := this.Top()

	return topValue

}

func (this *KthLargest) Size() int {
	return len(this.data)
}

func (this *KthLargest) Push(value int) {
	/**
	1. insert data tail
	2. adjustedUpward
	*/

	if this.Size() < this.cap {
		this.data = append(this.data, value)
		cur := len(this.data) - 1
		father := (cur - 1) / 2

		for father >= 0 && this.data[father] > this.data[cur] {
			this.data[father], this.data[cur] = this.data[cur], this.data[father]
			cur = father
			father = (cur - 1) / 2
		}
	} else {

		//if value >
		topValue, _ := this.Top()
		if value > topValue {
			this.data[0] = value
		}
		this.update(0, this.cap)
	}

}

func (this *KthLargest) Top() (int, error) {
	if this.Empty() {
		return -1, errors.New("heap empty")
	}
	return this.data[0], nil
}

func (this *KthLargest) update(pos int, n int) {
	if pos >= n-1 {
		return
	}

	lchild := pos*2 + 1
	rchild := pos*2 + 2

	min := pos

	if lchild < n && this.data[lchild] < this.data[min] {
		min = lchild
	}

	if rchild < n && this.data[rchild] < this.data[min] {
		min = rchild
	}

	if min != pos {
		this.data[min], this.data[pos] = this.data[pos], this.data[min]
		this.update(min, n)
	}

}

func (this *KthLargest) Empty() bool {
	return len(this.data) == 0
}
