package my_string

func MakeSmallestPalindrome(s string) string {
	length := len(s)
	bytes := []byte(s)
	leftIndex := 0
	rightIndex := length - 1
	for leftIndex < rightIndex {
		if bytes[leftIndex] != bytes[rightIndex] {
			if bytes[leftIndex] < bytes[rightIndex] {
				bytes[rightIndex] = bytes[leftIndex]
			} else {
				bytes[leftIndex] = bytes[rightIndex]
			}
		}
		leftIndex++
		rightIndex--
	}
	return string(bytes)
}
