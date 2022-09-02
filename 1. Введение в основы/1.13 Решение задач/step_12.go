// Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.


package main

import . "fmt"

func main() {
	var n int
	Scan(&n)

	var endString string

	n1, n2 := n%10, n%100

	if n1 == 1 && n2 != 11 {
		endString = "a"
	} else if (n1 == 2 && n2 != 12) || (n1 == 3 && n2 != 13) || (n1 == 4 && n2 != 14) {
		endString = "y"
	}

	Printf("%d korov%s", n, endString)
}