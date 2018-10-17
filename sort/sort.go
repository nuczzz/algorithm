package sort

type SortInterface interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

func Sort(data SortInterface) {
	InsertionSort(data)
	//BubbleSort(data)
	//CocktailSort(data)
	//SelectionSort(data)
}
