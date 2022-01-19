package section2

import "fmt"

func DemoAppend() {
	s := []int{1, 2}
	println(&s)
	s = appendInt(s, 2)
	s2 := []int{1, 2, 3}
	s = s2
	println(&s)
	println(&s2)
	fmt.Println(s)
}

func appendInt(s []int, n int) []int {
	var z []int
	//如果容量足够，直接推入
	//容量不够，扩展容量再推入
	sLen := len(s)
	zLen := sLen + 1

	if zLen <= cap(s) {
		z = s[:]
	} else {
		zCap := zLen
		if zCap < 2*sLen {
			zCap = 2 * sLen
		}
		z = make([]int, zLen, zCap)
		copy(z, s)
	}
	z[sLen] = n
	return z
}
