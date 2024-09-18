package main

import (
	"fmt"
	"os"
)

func main() {

	// Написать программу, которая считывает массив целых чисел и выводит их в обратном порядке.

	// создание переменных
	var numbers []int
	n := 0
	number := 0

	// получение размера массива
	fmt.Println("введите размер массива:")
	fmt.Fscan(os.Stdin, &n)

	// заполнение массива
	fmt.Println("заполните массив:")

	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &number)
		numbers = append(numbers, number)
	}

	// вывод изначального массива
	fmt.Println("было:\n", numbers)

	// переворачивание массива
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	// вывод результата
	fmt.Println("стало:\n", numbers)

}
