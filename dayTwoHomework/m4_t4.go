package main

import (
	"fmt"
	"log"
)

/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами

Пример:
Вход: 5
Выход:
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
*/

func checkBottlesCount(count int) {
	if count < 1 {
		log.Fatal("Количество должно быть больше 0!")
	}
}

func main() {
	var cn int
	fmt.Print("Введите количество: ")
	fmt.Scanf("%d", &cn)
	checkBottlesCount(cn)

	slice := make([][]int, cn)
	for i := range slice {
		slice[i] = make([]int, cn)
	}

	for i := range slice {
		for j := range slice {
			if (j+i)%2 != 0 {
				slice[i][j] = 1
			}
		}
	}

	for i := range slice {
		for j := range slice {
			fmt.Print(slice[i][j], " ")
		}
		fmt.Println()
	}
}
