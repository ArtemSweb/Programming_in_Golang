// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		floatRes := math.Pow(2, float64(i))
		if int(floatRes) <= n {
			fmt.Print(floatRes, " ")
		} else {
			return
		}
	}
}



