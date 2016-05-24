package sort

import (
	"testing"
	"fmt"
)

func Test_SelectSort(t *testing.T) {
	a := []int{2, 1, 3, 5, 4}
	a = SelectSort(a)

	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			t.Errorf("sort error")
		}
	}

}

func Test_PopSort(t *testing.T) {
	a := []int{2, 1, 3, 5, 4}
	a = popSort(a)

	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			t.Errorf("sort error")
		}
	}

}

func Test_QuickSort(t *testing.T) {
	a := []int{2, 1, 3, 5, 4}
	a = quickSort(a)

	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			t.Errorf("sort error")
		}
	}

}
