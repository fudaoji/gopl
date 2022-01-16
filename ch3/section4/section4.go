package section4

func Excer1() {

}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func Demo() {
	println(comma("121311"))
}
