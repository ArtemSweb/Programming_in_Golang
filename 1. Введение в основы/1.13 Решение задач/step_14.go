// Номер числа Фибоначчи

package main

import "fmt"

func main() {
	var n, i int

	fmt.Scan(&n)

	a := 0
	b := 1
	for i = 0; a < n; i++ {
		a, b = b, b+a
	}
	if a == n {
		fmt.Print(i)
	} else {
		fmt.Println(-1)
	}
}
