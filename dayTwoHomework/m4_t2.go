package main

import (
	"fmt"
	"log"
)

/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок

In: 1
Out: 1 бутылка

In: 22
Out: 22 бутылки
*/

func checkBottlesCount(count int) {
	if count < 0 || count > 200 {
		log.Fatal("Количество бутылок должно быть от 0 до 200!")
	}
}

func bottlesCountToTxt(count int) string {
	cn := count

	if cn > 19 {
		cn %= 10
	}

	if (cn) == 1 {
		return "бутылка"
	} else if cn > 1 && cn < 5 {
		return "бутылки"
	} else {
		return "бутылкок"
	}

}

func main() {
	var cn int
	fmt.Print("Введите количество бутылок: ")
	fmt.Scanf("%d", &cn)
	checkBottlesCount(cn)
	fmt.Println(cn, bottlesCountToTxt(cn))

}
