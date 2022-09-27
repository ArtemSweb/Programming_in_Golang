/*
Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X. Если подстроки S нет в строке X - вывести -1
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string
	var sub_s string
	fmt.Scan(&s, &sub_s)
	fmt.Println(strings.Index(s, sub_s))
}