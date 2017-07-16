package converter

type ConverterQueueItem struct {
	converter Converter
	priority  int
	index     int
}

type ConverterQueue []*ConverterQueueItem

func (c *ConverterQueueItem) Converter() Converter {
	return c.converter
}

func (pq ConverterQueue) Len() int { return len(pq) }

func (pq ConverterQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq ConverterQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *ConverterQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*ConverterQueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *ConverterQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
