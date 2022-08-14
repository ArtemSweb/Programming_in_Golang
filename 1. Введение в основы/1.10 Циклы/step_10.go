// Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа


package main

import "fmt"

func main() {
	var n1, n2 string

	fmt.Scan(&n1, &n2)

	for _, a1 := range n1 {
		for _, b1 := range n2 {
			if a1 == b1 {
				fmt.Print(string(a1) + " ")
			}
		}
	}
}