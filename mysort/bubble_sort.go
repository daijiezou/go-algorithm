package mysort

import "fmt"

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

func BubbleSort2(arrayList []int) {
	n := len(arrayList)
	for i := 0; i < n; i++ {
		swapFlag := false
		for j := n - 1; j > i; j-- {
			if arrayList[j] < arrayList[j-1] {
				arrayList[j], arrayList[j-1] = arrayList[j-1], arrayList[j]
				swapFlag = true
			}
		}
		// 一次交换操作都没有进行
		if !swapFlag {
			break
		}
	}
	fmt.Println(arrayList)
}

func Insert(arrayList []int) {
	n := len(arrayList)
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if arrayList[j] < arrayList[j-1] {
				arrayList[j], arrayList[j-1] = arrayList[j-1], arrayList[j]
			} else {
				break
			}
		}
	}
	fmt.Println(arrayList)
}
