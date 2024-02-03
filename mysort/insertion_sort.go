package mysort

func InsertionSort(arrayList []int) {
	length := len(arrayList)
	if length < 1 {
		return
	}
	for i := 1; i < length; i++ {
		value := arrayList[i]
		j := i - 1
		for ; j >= 0; j-- {
			// 最后一位数据比要插入的数据大，则移动
			if arrayList[j] > value {
				// 移动数据
				arrayList[j+1] = arrayList[j]
			} else {
				break
			}
		}
		// 插入数据
		arrayList[j+1] = value
	}
}

func InsertionSort2(arrayList []int) {
	length := len(arrayList)
	if length <= 1 {
		return
	}
	for i := 1; i < length; i++ {
		j := i - 1
		insertValue := arrayList[i]
		for ; j >= 0; j-- {
			if arrayList[j] > insertValue {
				arrayList[j+1] = arrayList[j]
			} else {
				break
			}
		}
		// 插入数据
		arrayList[j+1] = insertValue
	}
}
