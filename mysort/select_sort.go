package mysort

func SelectSort(arrayList []int) {
	length := len(arrayList)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arrayList[j] <= arrayList[min] {
				min = j
			}
		}
		arrayList[i], arrayList[min] = arrayList[min], arrayList[i]
	}
}
