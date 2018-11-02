package sort

func MergeSort(data SortInterface, start, m, end int) {

}

// left: data[start:m-1], right: data[m, end]
// p2最终为end+1
// p1最大为p2，p1不可能越界
func zmerge(data []int, start, m, end int) {
	preP1, p1, p2 := start, start, m
	leftLength := m - start
	for pos := start - 1; leftLength > 0; pos++ {
		if p2 > end {
			if p1 < m {
				//impossible
			} else if pos+1 == p1 {
				break
			} else {
				if preP1 == p1 {
					data[pos+1], data[p1] = data[p1], data[pos+1]
					if pos+1 == preP1 {
						preP1++
					}

					if p1 < end {
						p1++
						preP1 = p1
					} else {
						//do nothing
					}
				} else { //preP1 != p1
					if pos+1 == preP1 {
						data[pos+1], data[p1] = data[p1], data[pos+1]
					} else {
						data[pos+1], data[preP1], data[p1] = data[p1], data[pos+1], data[preP1]
					}
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

		if data[p1] < data[p2] {
			leftLength--

			if p1 < m {
				p1++
			} else {
				if preP1 == p1 {
					data[pos+1], data[p1] = data[p1], data[pos+1]
					if p2-p1 == 1 {
						//do nothing
						if pos < m {
							p1 = m
						} else {
							p1 = pos + 1
						}
					} else {
						p1++
					}
					preP1 = p1
				} else {
					data[pos+1], data[preP1], data[p1] = data[p1], data[pos+1], data[preP1]
					preP1++
					if p2-p1 == 1 {
						//do nothing
					} else {
						p1++
					}
				}
			}
		} else {
			if p1 < m {
				data[p1], data[p2] = data[p2], data[p1]
				p1 = p2
				preP1 = p1
			} else {
				if preP1 == p1 {
					data[pos+1], data[p1], data[p2] = data[p2], data[pos+1], data[p1]
					p1 = p2
				} else {
					data[pos+1], data[preP1], data[p2] = data[p2], data[pos+1], data[preP1]
					if p1-preP1 == 1 {
						// do nothing
					} else {
						preP1++
					}
				}
			}

			p2++
		}
	}
}

