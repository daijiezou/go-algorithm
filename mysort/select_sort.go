package mysort

func SelectSort(arrayList []int) {
	length := len(arrayList)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arrayList[j] <= arrayList[minIndex] {
				minIndex = j
			}
		}
		arrayList[i], arrayList[minIndex] = arrayList[minIndex], arrayList[i]
	}
}
