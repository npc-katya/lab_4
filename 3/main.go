package main

import (
	"fmt"
	"os"
)

func main() {

	// Написать программу, которая удаляет запись из карты по заданному имени.

	// создание и заполнение карты
	ages := map[string]int{
		"Eve":   13,
		"Sima":  23,
		"Alice": 25,
	}

	fmt.Println("введите одно из перечисленных имён: ")

	// вывод всех ключей
	for key := range ages {
		fmt.Print(key, " ")
	}
	fmt.Println()

	// получение ключа
	var str string
	fmt.Fscan(os.Stdin, &str)

	fmt.Println()

	// проверка существования ключа
	i := 0
	for key := range ages {
		if key == str {
			i++
		}
	}

	if i == 0 {
		fmt.Println("запись", str, "не найдена")
	} else {
		delete(ages, str)
		fmt.Println("запись", str, "удалена")
	}

	// вывод всех записей
	fmt.Println("\nвсе записи:")
	for key, i := range ages {
		fmt.Println(key, i)
	}

}
