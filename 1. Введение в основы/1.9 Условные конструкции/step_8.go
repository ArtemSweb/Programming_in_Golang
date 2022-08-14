package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fnum := num / 1000
	snum := num % 1000

	if fnum/100+(fnum/10)%10+fnum%10 == snum/100+(snum/10)%10+snum%10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}