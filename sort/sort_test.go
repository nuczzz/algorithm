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
