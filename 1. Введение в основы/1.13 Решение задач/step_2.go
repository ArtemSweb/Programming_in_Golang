// Дано трехзначное число. Найдите сумму его цифр. 


package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	fmt.Println((n / 100 % 10) + (n / 10 % 10) + (n % 10))
}