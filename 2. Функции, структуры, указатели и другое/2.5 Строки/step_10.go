/*
На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	rs := []rune(s)

	for i := range rs {
		if i%2 == 0 {
			rs[i] = ' '
		}
	}

	fmt.Println(strings.Replace(string(rs), " ", "", -1))

}