package sort


//快速排序是使用分治的思想进行排序的， 思路是选定一个基准值，然后将小于该值的数字放在左边
//将大于该值数字放在右边，接下来就可以一直这样迭代
//具体实现的逻辑就是：
//选取数组的最后一个值作为基准值， 同时从0开始设置两个值， i,j  使用j进行遍历， i指向的数字就是分界点，
//所以在正常的过程中，i指向的那个值是比基准值大， 在执行完一次排序后， i指向的值一定比基准值大， 所以在一次循环完
//的时候将i执行的基准值和最后一个值进行替换

func quickSort(input []int) []int {
	a:= partition(input)
	return a
}

func partition(a []int)( []int) {
	if len(a) <=1{
		return a
	}
	i:=0
	r:= len(a) -1
	for j:=0; j<len(a)-1; j++ {
		//至于j指向的值小于基准值的时候， 才将ij交换， 从而确保i指向的始终是大于基准值的，小于i
		//的都是小于基准值的
		if a[j]< a[r] {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]

	left := partition(a[:i])
	right := partition(a[i:])
	return append(left, right...)
}
