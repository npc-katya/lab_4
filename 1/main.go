package main

import (
	"fmt"
	"os"
)

func main() {

	// Написать программу, которая создает карту с именами людей и их возрастами. Добавить нового человека и вывести все записи на экран.

	// создание и заполнение карты
	ages := map[string]int{
		"Eve":   13,
		"Sima":  23,
		"Alice": 25,
	}

	// получение имени и возраста
	var str string
	var age int

	fmt.Println("введите имя и возраст:")

	fmt.Fscan(os.Stdin, &str)
	fmt.Fscan(os.Stdin, &age)

	// добавление записи в карту
	ages[str] = age

	// вывод всех записей
	fmt.Println("\nвсе записи:")
	for key, i := range ages {
		fmt.Println(key, i)
	}

}
