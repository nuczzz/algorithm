package sort

func CocktailSort(data SortInterface) {
	if data.Len() > 1 {
		left := 0
		right := data.Len() - 1
		for left < right {
			//right forward sort
			for i := left; i < right; i++ {
				if !data.Less(i, i+1) {
					data.Swap(i, i+1)
				}
			}
			right--

			//left forward sort
			for j := right; j > left; j-- {
				if data.Less(j, j-1) {
					data.Swap(j, j-1)
				}
			}
			left++
		}
	}
}