package sort

func MergeSort(data SortInterface, start, m, end int) {

}

// left: data[start:m-1], right: data[m, end]
// p2最终为end+1
// p1最大为p2，p1不可能越界
func merge(data []int, start, m, end int) {
	p1, p2 := start, m
	leftLength := m - start
	flag := false
	for pos := start - 1; pos < end; pos++ { //pos：已排序好的最后一个元素的数组下标不需要到end，因为最后剩一个元素时不需要操作
		//p2走完
		if p2 > end { //数组2数据已全部排序完，p1不可能在左侧
			if p2-p1 == 1 { //p1与p2相邻
				if p1-pos > 1 { //p1与pos之间还有数组1的数据
					data[pos+1], data[p1] = data[p1], data[pos+1] //交换两个数据位置，p1不变
					leftLength--
					//p1走完，不用管p2位置，排序完成
					if leftLength == 0 { //todo：升序左侧走完，降序右侧走完
						break
					}
				} else {
					break //左边数据已排序好，右边数据不需再处理
				}
			} else if p2-p1 > 1 { //p1~p2间还有数组1的数据
				if p1-pos <= 1 {
					break
				}

				data[pos+1], data[p1] = data[p1], data[pos+1] //先交换p1和需要处理数据的位置
				p1++                                          //p1后移
				leftLength--
				//p1走完，不用管p2位置，排序完成
				if leftLength == 0 { //todo：升序左侧走完，降序右侧走完
					break
				}
			} else {
				break // p1==p2，不需再处理
			}
			continue
		}

		if data[p1] < data[p2] { //取左边数组数据
			//todo:serious bug
			if p1 < m { //p1还在左侧数组范围内
				p1++
				leftLength--
				//p1走完，不用管p2位置，排序完成
				if leftLength == 0 { //todo：升序左侧走完，降序右侧走完
					break
				}
				//fmt.Println(data)
			} else { //p1已经到了右侧数组范围
				if p2-p1 == 1 { //p1与p2相邻
					if p1-pos > 1 { //p1与pos之间还有数组1的数据
						data[pos+1], data[p1] = data[p1], data[pos+1] //交换两个数据位置，p1不变
						leftLength--
						//p1走完，不用管p2位置，排序完成
						if leftLength == 0 { //todo：升序左侧走完，降序右侧走完
							break
						}
					} else {
						break //左边数据已排序好，右边数据不需再处理
					}
				} else if p2-p1 > 1 { //p1~p2间还有数组1的数据
					data[pos+1], data[p1] = data[p1], data[pos+1] //先交换p1和需要处理数据的位置
					p1++                                          //p1后移
					leftLength--
					//p1走完，不用管p2位置，排序完成
					if leftLength == 0 { //todo：升序左侧走完，降序右侧走完
						break
					}
				} else {
					break // p1==p2，不需再处理
				}
			}
		} else { //取右边数组数据
			data[pos+1], data[p2] = data[p2], data[pos+1] //交换两个数据位置
			if !flag {
				p1 = p2 //p1游标总是指向数组1中最小的数据，p1从左侧移到右侧只会发生一次
				flag = true
			}
			p2++ //p2游标右移
		}
	}
}

