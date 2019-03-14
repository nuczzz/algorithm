package floyd

import (
	"testing"
	"fmt"
)

func TestFloyd(t *testing.T) {
	const Infinity = -1
	matrix := [][]int{
		{Infinity, 12, Infinity, Infinity, Infinity, 16, 14},
		{12, Infinity, 10, Infinity, Infinity, 7, Infinity},
		{Infinity, 10, Infinity, 3, 5, 6, Infinity},
		{Infinity, Infinity, 3, Infinity, 4, Infinity, Infinity},
		{Infinity, Infinity, 5, 4, Infinity, 2, 8},
		{16, 7, 6, Infinity, 2, Infinity, 9},
		{14, Infinity, Infinity, Infinity, 8, 9, Infinity},
	}
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
}
