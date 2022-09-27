/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var pass string

	fmt.Scan(&pass)

	if isValid(pass) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func isValid(pass string) bool {
	rs := []rune(pass)

	if len(rs) < 5 {
		return false
	}
	for _, i := range rs {
		if !unicode.Is(unicode.Latin, i) && !unicode.Is(unicode.Digit, i) {
			return false
		}
	}
	return true
}