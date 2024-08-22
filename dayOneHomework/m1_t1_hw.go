package main

/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/

import "fmt"

const price = 48.0

func main() {
	var dist, exp int
	fmt.Print("Введите расторяние: ")
	fmt.Scanf("%d\n", &exp)
	if exp < 50 || exp > 10000 {
		fmt.Println("Расстояние должно быть от 50 до 10000км")
		return
	}
	fmt.Print("Введите расход топлива: ")
	fmt.Scanf("%d\n", &dist)
	if dist < 5 || dist > 25 {
		fmt.Println("Расход должен быть от 5 до 25л")
		return
	}

	fmt.Printf("Стоимость поездки в рублях: %.2f\n", float64(exp)/100*float64(dist)*price)
}
