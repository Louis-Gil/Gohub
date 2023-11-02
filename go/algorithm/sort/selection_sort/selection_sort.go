package myselectionsort

func MyselectionSort(s []int) {
	for i := 0; i < len(s) - 1; i++ {
		minIndex := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[minIndex] {
				minIndex = j
			}
		}
		s[i], s[minIndex] = s[minIndex], s[i]
	}
}
