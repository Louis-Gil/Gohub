package mybubblesort

import (
	"testing"
)

func TestMyBubbleSort(t *testing.T) {
	slice1 := []int{5, 3, 8, 4, 2}
	MyBubbleSort(slice1)

	if slice1[0] != 2 || slice1[1] != 3 || slice1[2] != 4 || slice1[3] != 5 || slice1[4] != 8 {
		t.Errorf("MyBubbleSort() = %v; want [2 3 4 5 8]", slice1)
	}

	slice2 := []int{1, 2, 3, 4, 5}
	MyBubbleSort(slice2)

	if slice2[0] != 1 || slice2[1] != 2 || slice2[2] != 3 || slice2[3] != 4 || slice2[4] != 5 {
		t.Errorf("MyBubbleSort() = %v; want [1 2 3 4 5]", slice2)
	}

	slice3 := []int{5, 4, 3, 2, 1}
	MyBubbleSort(slice3)

	if slice3[0] != 1 || slice3[1] != 2 || slice3[2] != 3 || slice3[3] != 4 || slice3[4] != 5 {
		t.Errorf("MyBubbleSort() = %v; want [1 2 3 4 5]", slice3)
	}

	slice4 := []int{5}
	MyBubbleSort(slice4)

	if slice4[0] != 5 {
		t.Errorf("MyBubbleSort() = %v; want [5]", slice4)
	}

	slice5 := []int{}
	MyBubbleSort(slice5)

	if len(slice5) != 0 {
		t.Errorf("MyBubbleSort() = %v; want []", slice5)
	}

	slice6 := []int{2, 3, 3, 4, 4, 5}
	MyBubbleSort(slice6)

	if slice6[0] != 2 || slice6[1] != 3 || slice6[2] != 3 || slice6[3] != 4 || slice6[4] != 4 || slice6[5] != 5 {
		t.Errorf("MyBubbleSort() = %v; want [2 3 3 4 4 5]", slice6)
	}

	slice7 := []int{5, 4, 3, 2, 1, 0, -1, -2, -3, -4}
	MyBubbleSort(slice7)

	if slice7[0] != -4 || slice7[1] != -3 || slice7[2] != -2 || slice7[3] != -1 || slice7[4] != 0 || slice7[5] != 1 || slice7[6] != 2 || slice7[7] != 3 || slice7[8] != 4 || slice7[9] != 5 {
		t.Errorf("MyBubbleSort() = %v; want [-4 -3 -2 -1 0 1 2 3 4 5]", slice7)
	}

	slice8 := []int{1358973, 3158712, 469024, 328095, 12, 3579, 4583, 12590, 494582, 237120, 125875130, 2119, 135900, 1298043, 94863427, 464896, 34612984, 3758192, 4609328357, 1328795}
	MyBubbleSort(slice8)

	if slice8[0] != 12 || slice8[1] != 2119 || slice8[2] != 3579 || slice8[3] != 4583 || slice8[4] != 12590 || slice8[5] != 135900 || slice8[6] != 237120 || slice8[7] != 328095 || slice8[8] != 464896 || slice8[9] != 469024 || slice8[10] != 494582 || slice8[11] != 1298043 || slice8[12] != 1328795 || slice8[13] != 1358973 || slice8[14] != 3158712 || slice8[15] != 3758192 || slice8[16] != 34612984 || slice8[17] != 94863427 || slice8[18] != 125875130 || slice8[19] != 4609328357 {
		t.Errorf("MyBubbleSort() = %v; want [12 2119 3579 4583 12590 135900 237120 328095 464896 469024 494582 1298043 1328795 1358973 3158712 3758192 34612984 94863427 125875130 4609328357]", slice8)
	}
}
