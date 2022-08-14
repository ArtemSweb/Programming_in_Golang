package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num/100 == (num/10)%10 || num/100 == num%10 || (num/10)%10 == num%10 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}