// Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.

package main

import "fmt"

func main() {
	var (
		n int
		c int
		d int
	)
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)

	for i := 0; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}

}