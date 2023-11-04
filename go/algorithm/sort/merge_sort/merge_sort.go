package mymergesort

func MyMergeSort(s []int) {
	result := mergeSort(s)
	copy(s, result)
}

func mergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	mid := len(s) / 2
	left := mergeSort(s[:mid])
	right := mergeSort(s[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	for i, j, k := 0, 0, 0; k < len(result); k++ {
		if i >= len(left) {
			result[k] = right[j]
			j++
		} else if j >= len(right) {
			result[k] = left[i]
			i++
		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}

	return result
}
