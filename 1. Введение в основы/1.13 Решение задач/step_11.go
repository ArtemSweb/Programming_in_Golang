// Найдите самое большее число на отрезке от a до b, кратное 7

package main

import . "fmt"

func main() {
	var a, b int
	Scan(&a, &b)

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			Print(i)
			return
		}
	}
	Print("NO")
}
