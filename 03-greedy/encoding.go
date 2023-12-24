package greedy

import (
	"github.com/calcifer777/learn-algorithms/data"
)

func Huffman(freqs map[string]int) data.BTree[string] {
	pq := data.NewPriorityQueue[data.BTree[string]](len(freqs))
	for s, freq := range freqs {
		pq.Push(data.NewBTreeNode(s, nil, nil), freq)
	}
	for {
		t1, freq1, _ := pq.Pop()
		t2, freq2, ok2 := pq.Pop()
		if !ok2 {
			return *t1
		}
		tMrg := data.NewBTreeNode(
			t1.Value()+t2.Value(),
			t1,
			t2,
		)
		pq.Push(tMrg, freq1+freq2)
	}
}
