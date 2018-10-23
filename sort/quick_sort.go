package sort

import (
	"fmt"
	"sync"
)

// when the length of data > 100, start goroutine to sort it
const LengthStartGoroutine = 100

// QuickSort sort data[start...end], contains data[end]
func QuickSort(data SortInterface, start, end int) {
	if start < end {
		q := partition(data, start, end)
		QuickSort(data, start, q-1)
		QuickSort(data, q+1, end)
	}
}

func partition(data SortInterface, start, end int) int {
	smallCount := start - 1
	for i := start; i < end; i++ {
		if data.Less(i, end) {
			smallCount++
			if smallCount != i {
				data.Swap(smallCount, i)
				fmt.Println(data)
			}
		}
	}

	if smallCount+1 != end {
		data.Swap(smallCount+1, end)
		fmt.Println(data)
	}

	return smallCount + 1
}

func QuickSortGoroutine(data SortInterface, start, end int, wg *sync.WaitGroup) {
	defer wg.Done()

	if start < end {
		q := partition(data, start, end)

		// todo: optimized by left and right data length
		newWg := &sync.WaitGroup{}
		newWg.Add(2)
		go QuickSortGoroutine(data, start, q-1, newWg)
		go QuickSortGoroutine(data, q+1, end, newWg)
		newWg.Wait()
	}
}
