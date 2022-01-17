package section6

import "math"

const a = iota

const (
	b = iota
	c
	d
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
)

func Excer1() {
	const (
		_ = math.Pow(1000, iota)
		KB
		MB
		GB
	)

	println(KB, MB, GB)
}

func Demo() {
	println(a, b, c, d)
	println(KiB, MiB, GiB)
}
