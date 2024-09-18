package main

import "fmt"

func main() {

	// Реализовать функцию, которая принимает карту и возвращает средний возраст всех людей в карте.

	// создание и заполнение карты
	ages := map[string]int{
		"Eve":   13,
		"Sima":  23,
		"Alice": 25,
	}

	// вывод всех возрастов
	fmt.Println("все возрасты:")

	for _, i := range ages {
		fmt.Print(i, " ")
	}

	// вывод среднего значение
	fmt.Println("\nсреднее значение:", avg(ages))

}

// функция, которая знаходит среднее значение всех возрастов из карты
func avg(ages map[string]int) (result float64) {
	k := 0
	i := 0

	for _, age := range ages {
		k = k + age
		i++
	}

	result = float64(k) / float64(i)

	return
}
