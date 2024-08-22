package main

import "fmt"

/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d\n", &num)

	if num < 1000 || num > 9999 {
		fmt.Println("Число должно быть четырёхзначным!")
		return
	}

	lf := num / 100
	rt := num % 100
	rt = rt%10*10 + rt/10

	if lf == rt {
		fmt.Printf("%d - палиндром", num)
	} else {
		fmt.Printf("%d - не палиндром", num)
	}
}
