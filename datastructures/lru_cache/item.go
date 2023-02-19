package main

type itemSlice []*item

type item struct {
	value interface{}
	age   int
}

func (items itemSlice) Len() int {
	return len(items)
}

func (items itemSlice) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items itemSlice) Less(i, j int) bool {
	return items[i].age < items[j].age
}
