package sort

func merge(data SortInterface, p, q, r int) {
	//todo
}

func MergeSort(data SortInterface, p, r int) {
	if p < r {
		q := (p+r)/2
		MergeSort(data, p, q)
		MergeSort(data, q+1, r)

		merge(data, p, q, r)
	}
}
