/*
На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input_text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input_text = strings.TrimSpace(input_text)
	text := []rune(input_text)

	var flag string = "Палиндром"

	for i := range text {
		for j := len(text) - 1; j >= 0; j-- {
			if string(text[i]) != string(text[j]) {
				flag = "Нет"
				break
			}
			break
		}
		break
	}
	fmt.Println(flag)

}
