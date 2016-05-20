package sort

//选择排序的golang实现
func SelectSort(input []int) []int {

	l := len(input)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	return input
}
