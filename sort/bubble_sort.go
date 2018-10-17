package sort

func BubbleSort(data SortInterface) {
	if data.Len() > 1 {
		for i := 0; i < data.Len(); i++ {
			for j := 1; j < data.Len()-i; j++ {
				if !data.Less(j-1, j) {
					data.Swap(j-1, j)
				}
			}
		}
	}
}
