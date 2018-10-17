package sort

func InsertionSort(data SortInterface) {
	if data.Len() > 1 {
		for j := 1; j < data.Len(); j++ {
			for i := j - 1; i >= 0; i-- {
				if !data.Less(i, i+1) {
					data.Swap(i, i+1)
				}
			}
		}
	}
}
