package sort

func SelectionSort(data SortInterface) {
	for i := 0; i < data.Len()-1; i++ {
		minIndex := i
		for j := i + 1; j < data.Len(); j++ {
			if !data.Less(minIndex, j) {
				minIndex = j
			}
		}

		if minIndex != i {
			data.Swap(i, minIndex)
		}
	}
}
