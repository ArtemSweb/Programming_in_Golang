package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	switch {
	case num < 10:
		fmt.Println(num)
	case 9 < num && num < 100:
		fmt.Println(num / 10)
	case 99 < num && num < 1000:
		fmt.Println(num / 100)
	case 999 < num && num < 10000:
		fmt.Println(num / 1000)
	case num >= 10000:
		fmt.Println(num / 10000)
	}
}