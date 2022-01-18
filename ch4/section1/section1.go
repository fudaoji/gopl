package section1

import (
	"crypto/sha256"
	"fmt"
)

func Demo() {
	arr := [...]string{0: "h", 1: "e"}

	fmt.Println(arr)
	fmt.Printf("%T\n", arr)
	arr2 := [...]int{99: -1}
	fmt.Println(arr2)
}

func Demosha256() {
	fmt.Printf("%T\n", sha256.Sum256([]byte("x")))
}

func SliceDemo() {
	s := make([]int, 3, 100)
	s1 := []int(nil)
	fmt.Println(s)
	fmt.Println(s1 == nil)
}
