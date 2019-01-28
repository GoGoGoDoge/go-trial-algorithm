package priorityqueue

import (
	"container/heap"
)

type Entry struct {
	key      string
	priority int
	// index is tricky, used when priority changes.
	index int
}

type PQ []*Entry

func (p PQ) Len() int { return len(p) }
func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}
func (p PQ) Less(i, j int) bool {
	// Pop higher priority first
	return p[i].priority > p[j].priority
}

func (p *PQ) Pop() interface{} {
	old := *p
	n := len(old)
	entry := old[n-1]
	entry.index = -1
	*p = old[0 : n-1]
	return entry
}

func (p *PQ) Push(x interface{}) {
	n := len(*p)
	entry := x.(*Entry)
	entry.index = n
	*p = append(*p, entry)
}

func (p *PQ) update(entry *Entry, key string, priority int) {
	entry.key = key
	entry.priority = priority
	heap.Fix(p, entry.index)
}
