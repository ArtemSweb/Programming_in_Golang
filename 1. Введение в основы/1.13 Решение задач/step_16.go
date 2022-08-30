// Из натурального числа удалить заданную цифру.


package main

import "fmt"

func main() {
	var s string
	var e string

	fmt.Scan(&s, &e)

	for i := 0; i < len(s); i++ {
		if string(s[i]) != e {
			fmt.Print(string(s[i]))
		}
	}
}