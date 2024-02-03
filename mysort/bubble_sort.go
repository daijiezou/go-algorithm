package mysort

// BubbleSort 冒泡排序
func BubbleSort(arrayList []int) {
	length := len(arrayList)
	for i := 0; i < length; i++ {
		// 提前退出冒泡循环的标志位
		var flag bool
		//
		for j := 0; j < length-i-1; j++ {
			if arrayList[j] > arrayList[j+1] {
				arrayList[j], arrayList[j+1] = arrayList[j+1], arrayList[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
