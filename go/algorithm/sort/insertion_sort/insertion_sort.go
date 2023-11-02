package myinsertionsort

func MyInsertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			} else {
				break
			}
		}
	}
}