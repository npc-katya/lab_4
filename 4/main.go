package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Написать программу, которая считывает строку с ввода и выводит её в верхнем регистре.

	var str string

	fmt.Println("введите строку:")

	// получение строки
	fmt.Fscan(os.Stdin, &str)

	// перевод строки в верхний регистр
	upperStr := strings.ToUpper(str)

	// вывод строки
	fmt.Println(upperStr)

}
