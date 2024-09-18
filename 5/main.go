package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Написать программу, которая считывает несколько чисел, введенных пользователем, и выводит их сумму.

	// создание сканера
	sc := bufio.NewScanner(os.Stdin)

	sum := 0

	// получение чисел
	fmt.Println("введите числа через пробел:")
	sc.Scan()
	a := sc.Text()

	// разделение строки на отдельные числа
	words := strings.Split(a, " ")

	for _, word := range words {
		num, _ := strconv.Atoi(word) // перевод из string в int
		sum += num                   // сложение
	}

	// вывод результата
	fmt.Println("сумма:", sum)

}
