package heap

import "errors"

type Heap struct {
	data []int
}

func (h *Heap) Push(value int) {
	/**
	1. insert data tail
	2. adjustedUpward
	*/
	h.data = append(h.data, value)
	cur := len(h.data) - 1
	father := (cur - 1) / 2

	for father >= 0 && h.data[father] > h.data[cur] {
		h.data[father], h.data[cur] = h.data[cur], h.data[father]
		cur = father
		father = (cur - 1) / 2
	}
}

func (h *Heap) Top() (int, error) {
	if h.Empty() {
		return -1, errors.New("heap empty")
	}
	return h.data[0], nil
}

func (h Heap) Pop() (int, error) {
	if h.Empty() {
		return -1, errors.New("heap empty")
	}

	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	h.data = h.data[0 : len(h.data)-1]
	// sinceDown
	h.update(0, len(h.data))
	return 0, nil
}

func (h Heap) update(pos int, n int) {
	if pos >= n-1 {
		return
	}

	lchild := pos*2 + 1
	rchild := pos*2 + 2

	min := pos

	if lchild < n && h.data[lchild] < h.data[min] {
		min = lchild
	}

	if rchild < n && h.data[rchild] < h.data[min] {
		min = rchild
	}

	if min != pos {
		h.data[min], h.data[pos] = h.data[pos], h.data[min]
		h.update(min, n)
	}

}

func (h *Heap) Empty() bool {
	return len(h.data) == 0
}

func (h *Heap) sort() {
	for i := len(h.data) - 1; i >= 1; i-- {
		h.data[i], h.data[0] = h.data[0], h.data[i]
		h.update(0, i)
	}
}
