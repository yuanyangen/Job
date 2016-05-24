package sort


func popSort(input []int) []int {
	l := len(input)

	for j := 0; j < l; j++ {
		for i := 0; i < l-j-1; i++ {
			if input[i] > input[i+1] {
				input[i+1], input[i] = input[i], input[i+1]
			}
		}
	}
	return input

}

