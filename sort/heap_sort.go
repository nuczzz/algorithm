package sort

func HeapSortSink(data SortInterface, i, length int) {
	left := i<<1+1
	right := i<<1+2

	var largest int
	if left < length && data.Less(i, left) {
		largest = left
	} else {
		largest = i
	}
	if right < length && data.Less(largest, right) {
		largest = right
	}

	if largest != i {
		data.Swap(i, largest)
		HeapSortSink(data, largest, length)
	}
}

func buildHeapSort(data SortInterface) {
	for i := data.Len()/2-1; i >= 0; i-- {
		HeapSortSink(data, i, data.Len())
	}
}

func HeapSort(data SortInterface) {
	if data.Len() > 1 {
		buildHeapSort(data)

		length := data.Len()
		for i := data.Len()-1; i > 0; i-- {
			data.Swap(i, 0)
			length--
			HeapSortSink(data, 0, length)
		}
	}
}
