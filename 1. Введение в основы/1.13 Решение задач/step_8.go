//По данным числам, определите количество чисел, которые равны нулю.


package main

import . "fmt"

func main() {
	var n int
	Scan(&n)
	cnt := 0
	lst := make([]int, n)

	for i := range lst {
		Scan(&lst[i])
	}

	for i := 0; i < len(lst); i++ {
		if lst[i] == 0 {
			cnt++
		}
	}

	Print(cnt)
}