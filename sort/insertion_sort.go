package sort

func InsertionSort(data SortInterface) {
	if data.Len() > 1 {
		for j := 1; j < data.Len(); j++ {
			for i := j - 1; i >= 0 && !data.Less(i, i+1); i-- {
				data.Swap(i, i+1)
			}
		}
	}
}
