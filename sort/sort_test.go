package sort

import (
	"testing"
	sysSort "sort"
)

type Array []int

func (a Array)Len() int {
	return len(a)
}

func (a Array)Less(i, j int) bool {
	return a[i] < a[j]
}

func (a Array)Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func TestSysSort(t *testing.T) {
	array := Array{2,3,4,6,5}
	sysSort.Sort(array)
	t.Log(array)
}

func TestMySort(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	Sort(data)
	t.Log(data)
}

func TestBubbleSort(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	BubbleSort(data)
	t.Log(data)
}

func TestCocktailSort(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	CocktailSort(data)
	t.Log(data)
}

func TestSelectionSort(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	SelectionSort(data)
	t.Log(data)
}

func TestBinaryInsertionSort(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	BinaryInsertionSort(data)
	t.Log(data)
}

func TestShellSort(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	ShellSort(data)
	t.Log(data)
}

func TestQuickSort(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	QuickSort(data, 0, data.Len()-1)
	t.Log(data)
}

func TestQuickSortGoroutine(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	QuickSortGoroutine(data, 0, data.Len()-1, nil)
	t.Log(data)
}

func TestMergeSort(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	MergeSort(data, 0, data.Len()-1)
	t.Log(data)
}

func TestHeapSort(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	HeapSort(data)
	t.Log(data)
}
