package sort

func InsertionSort(data SortInterface) {
	for j := 1; j < data.Len(); j++ {
		for i := j - 1; i >= 0; i-- {
			if !data.Less(i, i+1) {
				data.Swap(i, i+1)
			}
		}
	}
}

func BinaryInsertionSort(data SortInterface) {
	for j := 1; j < data.Len(); j++ {
		left, middle, right := 0, 0, j-1
		for left < right {
			middle = (left + right) / 2
			if data.Less(middle, j) {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}

		for i := j; i > left; i-- {
			if data.Less(i, i-1) {
				data.Swap(i, i-1)
			}
		}
	}
}

func ShellSort(data SortInterface) {
	for gap := data.Len() / 2; gap > 0; gap /= 2 {
		for i := gap; i < data.Len(); i++ {
			for j := i; j-gap >= 0; j -= gap {
				if data.Less(j, j-gap) {
					data.Swap(j-gap, j)
				}
			}
		}
	}
}
