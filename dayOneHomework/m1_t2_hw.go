package main

import "fmt"

/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d\n", &num)

	if num < 100 || num > 999 {
		fmt.Println("Число должно быть трёхзначным!")
		return
	}

	first_digit := num / 100
	second_digit := num / 10 % 10
	third_digit := num % 10

	fmt.Printf("Реверсная запись: %d%d%d", third_digit, second_digit, first_digit)
}
