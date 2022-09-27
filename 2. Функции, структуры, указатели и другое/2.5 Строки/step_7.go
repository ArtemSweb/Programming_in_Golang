/*
На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	input_text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input_text = strings.TrimSpace(input_text)
	text := []rune(input_text)

	if strings.HasSuffix(input_text, ".") && unicode.IsUpper(text[0]) {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}