package main

import (
	"container/heap"
	"fmt"
)

type Entry struct {
	key      string
	priority int
	// index is tricky, used when priority changes.
	index int
}

type pq []*Entry

func (p pq) Len() int { return len(p) }
func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}
func (p pq) Less(i, j int) bool {
	// Pop higher priority first
	return p[i].priority > p[j].priority
}

func (p *pq) Pop() interface{} {
	old := *p
	n := len(old)
	entry := old[n-1]
	entry.index = -1
	*p = old[0 : n-1]
	return entry
}

func (p *pq) Push(x interface{}) {
	n := len(*p)
	entry := x.(*Entry)
	entry.index = n
	*p = append(*p, entry)
}

func (p *pq) update(entry *Entry, key string, priority int) {
	entry.key = key
	entry.priority = priority
	heap.Fix(p, entry.index)
}

func main() {
	entries := map[string]int{
		"marco": 4, "maggie": 2, "alan": 1,
	}
	pq := make(pq, len(entries))
	i := 0
	for value, p := range entries {
		pq[i] = &Entry{
			key:      value,
			priority: p,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	newEntry := &Entry{
		key:      "Apple",
		priority: 0,
	}
	heap.Push(&pq, newEntry)
	pq.update(newEntry, "Apple", 5)
	for pq.Len() > 0 {
		fmt.Println(heap.Pop(&pq))
	}
}
