package mixed

import (
	"math"
	"fmt"
)

//这个函数的功能是求出输入的某个数的所有质数
//实现的方式是通过迭代的方式，从2开始，如果这个数是质数，设置其倍数也为质数
//这里包含了另外一个条件: 一个合数的所有的因子不会超过这个数开方


func Prime(num int) []int {
	result := []bool {} //true表示为是质数, 初始化的过程我也是醉了
	for i:=0; i<= num; i++ {
		result = append(result, true)
	}

	max := int(math.Sqrt(float64(num)))
	for i:=2; i<= max; i++ {
		if result[i] == true {
			for j:=2; j*i<=num; j++  {
				result[j*i] = false
			}
		}
	}
	var  ret []int
	for i,v := range result {
		if v == true && i!=0 {
			ret = append(ret, i)
		}
	}
	fmt.Println(ret)
	return ret
}

