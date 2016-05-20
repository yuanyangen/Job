package sort

import "testing"

func Test_SelectSort(t *testing.T) {
	a := []int{2, 1, 3, 5, 4}
	a = SelectSort(a)

	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			t.Errorf("sort error")
		}
	}

}
