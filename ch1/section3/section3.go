package section3

import (
	"bufio"
	"os"
)

//countLines  count line
func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}

//Demo2
func Demo2() {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)

	for line, n := range counts {
		if n > 1 {
			println(line, "=>", n)
		}
	}
}

//Demo
func Demo() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	for line, n := range counts {
		println(line, "=>", n)
	}
}
