package slice

import "testing"

func TestIArrayCopy(t *testing.T) {
	array := [][]int{{1,1,1}, {2,2,2}, {}}

	newArray := IArrayCopy(array)

	if len(array) != len(newArray) || len(array[0]) != len(newArray[0]){
		t.Fatalf("failed copy array, size error")
	}

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if array[i][j] != newArray[i][j] {
				t.Fatalf("failed copy array, num error")
			}
		}
	}
}
