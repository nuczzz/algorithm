package sort

import "fmt"

func MergeSort(data SortInterface, start, end int) {
	if start < end {
		m := (end + start) >> 1
		MergeSort(data, start, m)
		MergeSort(data, m+1, end)
		zMerge(data, start, m+1, end)
		fmt.Println(data)
	}
}

// left: data[start...m-1], right: data[m...end]
// todo: fix bug
func zMerge(data SortInterface, start, m, end int) {
	fmt.Println()
	fmt.Println(start)
	fmt.Println(m)
	fmt.Println(end)
	if m > end {
		return
	}

	preP1, p1, p2 := start, start, m
	leftLength := m - start
	//var count int
	for pos := start; leftLength > 0; pos++ {
		fmt.Println("*****")
		fmt.Println(pos)
		fmt.Println(preP1)
		fmt.Println(p1)
		fmt.Println(p2)
		fmt.Println(data)
		fmt.Println("****")
		fmt.Println()
		//count++
		if p2 > end {
			if pos == p1 {
				break
			} else {
				if preP1 == p1 {
					data.Swap(pos, p1)
					if p1 == end {
						p1 = pos+1
						preP1 = p1
					} else {
						p1++
						preP1 = p1
					}
				} else { //preP1 != p1
					data.Swap(pos, preP1)
					data.Swap(pos, p1)
					preP1++
					if p1 < end {
						p1++
					} else {
						p1 = preP1
					}
				}
			}
			continue
		}

		if data.Less(p1, p2) {
			leftLength--

			if p1 < m {
				p1++
				preP1 = p1
			} else {
				if preP1 == p1 {
					data.Swap(pos, p1)
					if pos < m {
						if p1+1 == p2 {
							p1 = m
						} else {
							p1++
						}
						preP1 = p1
					} else {
						if p1+1 < p2 {
							p1++
							//don't move preP1
						}
					}
				} else {
					if pos < m && preP1 > m {
						data.Swap(pos, preP1)
						data.Swap(pos, p1)
						preP1++
					} else {
						data.Swap(pos, p1)
					}

					if p1+1 < p2 {
						p1++
					} else {
						//do nothing
					}
				}
			}
		} else {
			if p1 < m {
				data.Swap(p1, p2)
				p1 = p2
				preP1 = p1
			} else {
				if preP1 == p1 {
					data.Swap(pos, p1)
					data.Swap(pos, p2)
					p1 = p2
					//don't move preP1
				} else {
					data.Swap(pos, preP1)
					data.Swap(pos, p2)
					//preP1++
				}
			}

			p2++
		}
	}
	//fmt.Println(count)
}
