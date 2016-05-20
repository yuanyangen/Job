package mixed

import (
	"strconv"
	"strings"
)

func Multiple(A string, B string) string {
	ret := doMultiple(A, B)
	trimed := ""

	for i := 0; i < len(ret); i++ {
		if string(ret[i]) == "0" {
			trimed += "0"
			continue
		}
		break
	}

	return strings.TrimPrefix(ret, trimed)

}

//实现了大数的乘法和加法，其中大数的加法是做乘法必须的一个部分
//实现的方法是A*B = (a*100 +b) *(c*100+d) = a*c*10000+(b*c+a*d)*100+b*d
func doMultiple(A string, B string) string {
	//将输入A, B转换成注释所示的公式
	var a string
	var b string
	var c string
	var d string
	//只有最终传入的数字是等于2的时候， 才进行最终的计算
	if len(A) <= 2 && len(B) <= 2 {
		tmpA, _ := strconv.Atoi(A)
		tmpB, _ := strconv.Atoi(B)
		ret := strconv.Itoa(tmpA * tmpB)
		return ret
	}

	//拆分出abcd四个个参数
	if len(A) > 2 {
		a = string(A[0:(len(A) - 2)])
		b = string(A[(len(A) - 2):])
	} else {
		a = "0"
		b = A
	}
	if len(B) > 2 {
		c = string(B[0:(len(B) - 2)])
		d = string(B[(len(B) - 2):])
	} else {
		c = "0"
		d = B
	}

	//组合形成公式
	ret := "0"
	ret = doMultiple(a, c) + "0000"
	ret = Add(ret, doMultiple(b, c)+"00")
	ret = Add(ret, doMultiple(a, d)+"00")
	ret = Add(ret, doMultiple(b, d))
	return ret
}

//实现大数的加法，在slice中，低位表示自然数中的低位
func Add(a string, b string) string {
	//将字符串的长度补充道有一致
	lenA := len(a)
	lenB := len(b)
	lent := 0
	if lenA < lenB {
		lent = lenB
		for i := 0; i < lenB-lenA; i++ {
			a = "0" + a
		}
	} else if lenB < lenA {
		lent = lenA
		for i := 0; i < lenA-lenB; i++ {
			b = "0" + b
		}
	} else {
		lent = lenA
	}

	ret := ""
	i := 1
	//进位标识
	flag := 0
	//迭代加法
	for i <= len(a) {
		tempA, _ := strconv.Atoi(string(a[lent-i]))
		tempB, _ := strconv.Atoi(string(b[lent-i]))

		//这两位相加的结果
		tmp := tempB + tempA + flag
		if tmp >= 10 {
			flag = 1
			tmp -= 10
		} else {
			flag = 0
		}
		ret = strconv.Itoa(tmp) + ret

		i++
	}
	//补充进位
	if flag == 1 {
		ret = "1" + ret
	}

	return ret

	//对字符串进行遍历，得到结果

}
