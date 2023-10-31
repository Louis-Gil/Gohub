package bubblesort

func MyBubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s) -i -1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}