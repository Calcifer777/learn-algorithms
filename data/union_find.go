package data

type UnionFind struct {
	records []Record
}

type Record struct {
	n             int
	component     *Record
	componentSize int
}

func NewUnionFind(nodes []int) UnionFind {
	// O(n)
	records := make([]Record, len(nodes))
	for i, n := range nodes {
		records[i] = Record{n, nil, 1}
	}
	return UnionFind{records: records}
}

func (uf *UnionFind) Union(s1, s2 int) {
	// O(1), no loops!
	c1 := uf.Find(s1)
	c2 := uf.Find(s2)
	var toChange int
	var masterRecord *Record
	var mergedRecordSize int
	size1 := uf.records[c1].componentSize
	size2 := uf.records[c2].componentSize
	if s1 > s2 {
		masterRecord = &uf.records[c1]
		toChange = s2
		mergedRecordSize = size2
	} else if uf.records[c1].componentSize < uf.records[c2].componentSize {
		masterRecord = &uf.records[c2]
		toChange = s1
		mergedRecordSize = size1
	} else if c1 < c2 {
		// in case of comp. length tie, merge into lower node
		masterRecord = &uf.records[c1]
		toChange = s2
		mergedRecordSize = size2
	} else {
		masterRecord = &uf.records[c2]
		toChange = s1
		mergedRecordSize = size1
	}
	masterRecord.componentSize += mergedRecordSize
	uf.records[toChange].component = masterRecord
}

func (uf *UnionFind) Find(n int) int {
	// O(log n): at every Union, component size doubles
	// 			 -> at most n links in the record.component chain
	for {
		c := uf.records[n].component
		if c == nil {
			return n
		} else {
			c := uf.Find(c.n)
			// update each record on the component chain with the master record
			uf.records[n].component = &uf.records[c]
			uf.records[n].componentSize = uf.records[c].componentSize
			return c
		}
	}
}
