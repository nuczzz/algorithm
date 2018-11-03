package sort

import (
	"testing"
	sysSort "sort"
	"math/rand"
	"time"
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

func TestQuickSortWithGoroutine(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	QuickSortWithGoroutine(data, 0, data.Len()-1, nil)
	t.Log(data)
}

func TestMergeSort(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	MergeSort(data, 0, data.Len()-1)
	t.Log(data)
}

func TestZMerge(t *testing.T) {
	data := Array{9, 31, 31, 37, 47, 0, 6, 18, 25, 40}
	zMerge(data, 0, 5, 9)
	t.Log(data)
}

func TestMergeAndQuickSort(t *testing.T) {
	num := 20
	data := make(Array, num)
	for i := 0; i < num; i++ {
		data[i] = rand.Intn(50)
	}
	t.Log(data)

	data1 := make(Array, num)
	data2 := make(Array, num)
	copy(data1, data)
	copy(data2, data)

	var now int64
	/*now = time.Now().UnixNano()
	t.Log(now)
	QuickSort(data1, 0, len(data)-1)
	now = time.Now().UnixNano()
	t.Log(now)
	//t.Log(data1)
	t.Log("\n")*/

	now = time.Now().UnixNano()
	t.Log(now)
	MergeSort(data2, 0, len(data)-1)
	now = time.Now().UnixNano()
	t.Log(now)
	t.Log(data2)
}

func TestHeapSort(t *testing.T) {
	data := Array{4, 6, 5, 3, 2, 1}
	HeapSort(data)
	t.Log(data)
}
