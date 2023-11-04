package myquicksort

func MyQuickSort(s []int) {
	quickSort(s, 0, len(s)-1)
}

func quickSort(s []int, low int, high int) {
	if low < high {
		pivotIndex := medianOfThree(s, low, high)
		newPivotIndex := partition(s, low, high, pivotIndex)
		quickSort(s, low, newPivotIndex-1)
		quickSort(s, newPivotIndex+1, high)
	}
}

func medianOfThree(s []int, low int, high int) int {
	mid := low + (high-low)/2
	if s[low] > s[mid] {
		s[low], s[mid] = s[mid], s[low]
	}
	if s[low] > s[high] {
		s[low], s[high] = s[high], s[low]
	}
	if s[mid] > s[high] {
		s[mid], s[high] = s[high], s[mid]
	}
	return mid
}

func partition(s []int, low int, high int, pivotIndex int) int {
	pivot := s[pivotIndex]
	s[pivotIndex], s[high] = s[high], s[pivotIndex]
	storeIndex := low
	for i := low; i < high; i++ {
		if s[i] < pivot {
			s[i], s[storeIndex] = s[storeIndex], s[i]
			storeIndex++
		}
	}
	s[storeIndex], s[high] = s[high], s[storeIndex]
	return storeIndex
}
