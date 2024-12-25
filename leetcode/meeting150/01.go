package meeting150

func merge(nums1 []int, m int, nums2 []int, n int) {
	mi, ni, p := m-1, n-1, len(nums1)-1
	for mi >= 0 && ni >= 0 {
		if nums1[mi] > nums2[ni] {
			nums1[p] = nums1[mi]
			p--
			mi--
		} else {
			nums1[p] = nums2[ni]
			p--
			ni--
		}
	}
	for mi >= 0 {
		nums1[p] = nums1[mi]
		p--
		mi--
	}
	for ni >= 0 {
		nums1[p] = nums2[ni]
		p--
		ni--
	}
}
