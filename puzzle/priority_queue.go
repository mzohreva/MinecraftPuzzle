package puzzle

import (
	"container/heap"
)

type minHeap struct {
	mh *minHeapImpl
}

func newMinHeap() minHeap {
	mh := make(minHeapImpl, 0)
	heap.Init(&mh)
	return minHeap{mh: &mh}
}

func (h minHeap) len() int      { return h.mh.Len() }
func (h minHeap) isEmpty() bool { return h.mh.Len() == 0 }

func (h minHeap) push(s State, path []Action, cost int) {
	node := &minHeapNode{cost: cost, state: s, path: path}
	heap.Push(h.mh, node)
}

func (h minHeap) pop() (State, []Action) {
	node := heap.Pop(h.mh).(*minHeapNode)
	return node.state, node.path
}

type minHeapNode struct {
	cost  int
	state State
	path  []Action
	index int
}

type minHeapImpl []*minHeapNode

func (mh minHeapImpl) Len() int           { return len(mh) }
func (mh minHeapImpl) Less(i, j int) bool { return mh[i].cost < mh[j].cost }
func (mh minHeapImpl) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
	mh[i].index = i
	mh[j].index = j
}

func (mh *minHeapImpl) Push(x interface{}) {
	n := len(*mh)
	item := x.(*minHeapNode)
	item.index = n
	*mh = append(*mh, item)
}

func (mh *minHeapImpl) Pop() interface{} {
	old := *mh
	n := len(old)
	item := old[n-1]
	item.index = -1
	*mh = old[0 : n-1]
	return item
}
